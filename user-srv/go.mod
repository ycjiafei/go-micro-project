module github.com/ycjiafei/go-micro-project/user-srv

go 1.14

require (
	github.com/golang/protobuf v1.4.1
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/ycjiafei/go-micro-project/database v0.0.0-20200820133556-262eb27eaef3
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
)

replace (
	github.com/ycjiafei/go-micro-project/database => ../database
)