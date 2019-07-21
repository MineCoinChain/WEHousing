package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"IHome/PutOrders/handler"
	example "IHome/PutOrders/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.PutOrders"),
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
