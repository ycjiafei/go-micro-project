.PHONY : proto-user
proto-user :
	protoc --go_out=plugins=grpc:. proto/user/user-srv.proto