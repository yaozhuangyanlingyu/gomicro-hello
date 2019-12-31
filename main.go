package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"aplum.com/hello/handler"
	"aplum.com/hello/subscriber"

	hello "aplum.com/hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hello", service.Server(), new(subscriber.Hello))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hello", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
