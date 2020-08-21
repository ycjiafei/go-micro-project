package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"net/http"
)

const (
	TracerCtx = "tracer_ctx"
	SpCtx = "span_ctx"
)

func JaegerTrace(c *gin.Context) {
	var sp opentracing.Span
	md := make(map[string]string)
	spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	// 两种初始化 span 的方式
	if err != nil {
		sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
	} else {
		sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
	}
	defer sp.Finish()
	// 记录 data 参数要放在 next 上面, 不然就被读走了
	sp.SetTag("http.data", getRequestData(c))

	opentracing.GlobalTracer().Inject(
		sp.Context(),
		opentracing.TextMap,
		opentracing.TextMapCarrier(md),
	)

	ctx := opentracing.ContextWithSpan(context.TODO(), sp)
	// 写入到整个请求上下文中
	c.Set(TracerCtx, ctx)
	c.Set(SpCtx, sp.Context())



	c.Next()


	// 记录这次 http 请求的信息
	statusCode := c.Writer.Status()
	if statusCode != http.StatusOK {
		ext.Error.Set(sp, true)
	}
	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
	ext.HTTPMethod.Set(sp, c.Request.Method)
	ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
	sp.SetTag("http.header", c.Request.Header)
	sp.SetTag("http.request_ip", c.ClientIP())
}

/**
由于 gin 的参数读一遍就消失了, 所以读完还要放回去
尝试过 c.Copy() 方法,  不知道为什么没有用
 */
func getRequestData(c *gin.Context) string {
	data, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return string(data)
}