package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"IHome/GetHouses/handler"
	example "IHome/GetHouses/proto/example"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.GetHouses"),
		micro.Version("latest"),
	)

	service.Init()

	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
