FROM golang:1.14 AS build-env
ADD . /go/src/app
WORKDIR /go/src/app
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux GOARCH=386 go build -v -o /go/src/app/user-server

FROM alpine
COPY --from=build-env /go/src/app/user-server /usr/local/bin/user-server
EXPOSE 50050
CMD [ "user-server" ]