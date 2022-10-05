// github.com/lcmaguire/protoc-gen-go-goo/example
package main

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleservice"
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
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	example.RegisterExampleServiceServer(server, &exampleservice.ExampleService{}) // this would need to be a list or multiple.
	reflection.Register(server)                                                    // this should perhaps be optional
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}
