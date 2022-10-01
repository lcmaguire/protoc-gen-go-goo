package main

import (
	log "log"
	net "net"

	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleservice"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	out.RegisterExampleServiceServer(server, &exampleservice.ExampleService{}) // this would need to be a list or multiple.
	reflection.Register(server)
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
