package main

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/example/exampleservice"
	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
	log "log"
	net "net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	listenOn := "127.0.0.1:8080"                 // this should be passed in via config
	listener, err := net.Listen("tcp", listenOn) // this too
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	// services in your protoFile

	example.RegisterExampleServiceServer(server, &exampleservice.ExampleService{})
	reflection.Register(server) // this should perhaps be optional

	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
