package main

import (
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/ycjiafei/go-micro-project/api/routes"
	slog "github.com/ycjiafei/go-micro-project/pkg/log"
	"github.com/ycjiafei/go-micro-project/pkg/trace"
	"go.uber.org/zap"
)

func main() {
	slog.InitDefault()
	tracer, closer, err := trace.NewJaeger("HTTP")
	if err != nil {
		slog.SLog.Bg().Error("初始化 jaeger 失败", zap.String("err", err.Error()))
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	routes.InitRoutes().Run()
}
