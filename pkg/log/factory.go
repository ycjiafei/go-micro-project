package log

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Factory struct {
	logger *zap.Logger
}

var SLog Factory

func InitDefault() {
	SLog = DefaultFactory()
}

func DefaultFactory() Factory {
	logger, _ := zap.NewDevelopment()
	return NewFactory(logger)
}

func NewFactory(logger *zap.Logger) Factory {
	return Factory{logger: logger}
}

/**
从 context 中取出 tracer 组件
 */
func (b Factory) For(ctx context.Context) Logger {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		logger := spanLogger{span: span, logger: b.logger}
		// 如果 tracer 是 jaeger 实现
		// 日志额外添加 traceID 和 spanID
		if jaegerCtx, ok := span.Context().(jaeger.SpanContext); ok {
			logger.spanFields = []zapcore.Field{
				zap.String("trace_id", jaegerCtx.TraceID().String()),
				zap.String("span_id", jaegerCtx.SpanID().String()),
			}
		}

		return logger
	}
	return b.Bg()
}

/**
将 factory 转换成 Logger 接口实现
 */
func (b Factory) Bg() Logger {
	return logger(b)
}

/**
派生 child logger
 */
func (b Factory) With(fields ...zapcore.Field) Factory {
	return Factory{
		logger: b.logger.With(fields...),
	}
}