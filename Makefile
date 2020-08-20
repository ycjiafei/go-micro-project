.PHONY : proto-user
	proto-user :
 	protoc --go_out=plugins=grpc:. user-srv/proto/user-srv.proto
.PHONY : start-api
start-api :
	go build