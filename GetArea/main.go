package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"IHome/GetArea/handler"

	example "IHome/GetArea/proto/example"
	"github.com/micro/go-grpc"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init()

	// 服务注册
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
