package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"IHome/GetUserInfo/handler"
	example "IHome/GetUserInfo/proto/example"
	"github.com/micro/go-grpc"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.GetUserInfo"),
		micro.Version("latest"),
	)

	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
