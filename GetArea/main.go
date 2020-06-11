package main

import (
	"GoMicroServer/GetArea/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.GetArea"),
		micro.Version("latest"),
		micro.Address(":8080"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	//GetArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), new(subscriber.GetArea))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.GetArea", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
