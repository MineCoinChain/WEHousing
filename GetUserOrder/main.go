package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"IHome/GetUserOrder/handler"
	example "IHome/GetUserOrder/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetUserOrder"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
