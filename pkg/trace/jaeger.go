package trace

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func NewJaeger(serviceName string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		// 设置采样率
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,	// 固定采样
			Param: 1,						// 1 全采样  0 不采样
		},

		Reporter: &config.ReporterConfig{
			LogSpans: true,
			BufferFlushInterval: time.Second,
			LocalAgentHostPort: os.Getenv("JAEGER_AGENT_HOST"),
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	return tracer, closer, err
}

type MDReaderWriter struct {
	metadata.MD
}

func (md MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vMD := range md.MD {
		for _, v := range vMD {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (md MDReaderWriter) Set(key string, val string) {
	key = strings.ToLower(key)
	md.MD[key] = append(md.MD[key], val)
}

func ClientInterceptor(tracer opentracing.Tracer, spCtx opentracing.SpanContext) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		sp := opentracing.SpanFromContext(ctx)
		if sp != nil {
			spCtx = sp.Context()
		}

		childSp := tracer.StartSpan(
			method,
			opentracing.ChildOf(spCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "GRPC CLIENT"},
			ext.SpanKindRPCClient,
		)
		defer childSp.Finish()
		// 记录下信息
		childSp.SetTag("client req", req)
		childSp.SetTag("server reply", reply)


		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		err := tracer.Inject(childSp.Context(), opentracing.TextMap, MDReaderWriter{md})
		if err != nil {
			log.Println("grpc client tracer inject err :", err)
		}

		newCtx := metadata.NewOutgoingContext(ctx, md)

		err = invoker(newCtx, method, req, reply, cc, opts...)
		if err != nil {
			log.Println("grpc client call err : ", err)
		}

		return err
	}
}

func ServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctxTracer, ok := ctx.Value("tracer_ctx").(opentracing.Tracer)
		// 如果 ctx 中有, 则替换调传入的
		if ok {
			tracer = ctxTracer
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		spCtx, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil {
			log.Println("grpc server extract err : ", err)
		} else {
			sp := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "GRPC"},
				ext.SpanKindRPCServer,
			)
			defer sp.Finish()
			sp.SetTag("received", req)
			ctx = opentracing.ContextWithSpan(ctx, sp)
		}
		return handler(ctx, req)
	}
}