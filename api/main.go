package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/ycjiafei/go-micro-project/api/routes"
	"github.com/ycjiafei/go-micro-project/pkg/trace"
	"log"
)

func main() {
	tracer, closer, err := trace.NewJaeger("HTTP")
	if err != nil {
		log.Printf("初始化 jaeger 失败 : %v \n", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	routes.InitRoutes().Run()
}
