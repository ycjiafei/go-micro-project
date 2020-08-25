module github.com/ycjiafei/go-micro-project/user-srv

go 1.14

require (
	github.com/golang/protobuf v1.4.1
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/ycjiafei/go-micro-project/database v0.0.0-20200820133556-262eb27eaef3
	github.com/ycjiafei/go-micro-project/pkg v0.0.0-20200824114841-804d97a18e5d // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0

)

// 此举是为了方便本地开发不用频繁去 github 更新
// 如果是部署 docker , 请注释掉
//replace (
//	github.com/ycjiafei/go-micro-project/database => ../database
//	github.com/ycjiafei/go-micro-project/pkg => ../pkg
//)
