module github.com/ycjiafei/go-micro-project/api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ycjiafei/go-micro-project/database v0.0.0-20200820133556-262eb27eaef3
	github.com/ycjiafei/go-micro-project/pkg v0.0.0-20200821042131-829ea6b31ed4
	github.com/ycjiafei/go-micro-project/user-srv v0.0.0-20200820125343-b0b51cbe68ca
	golang.org/x/net v0.0.0-20200813134508-3edf25e44fcc // indirect
	golang.org/x/sys v0.0.0-20200819171115-d785dc25833f // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.4.0
	google.golang.org/genproto v0.0.0-20200815001618-f69a88009b70 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)


// 此举是为了方便本地开发不用频繁去 github 更新
// 如果是部署 docker , 请注释掉
replace (
	github.com/ycjiafei/go-micro-project/database => ../database
	github.com/ycjiafei/go-micro-project/pkg => ../pkg
	github.com/ycjiafei/go-micro-project/user-srv => ../user-srv
)
