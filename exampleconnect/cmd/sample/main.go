package main

import (
	"log"
	"net/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	// your protoPathHere
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"

	// your services
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleservice"
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/streamingservice"
)

func main() {
	mux := http.NewServeMux()
	/*
		reflector := grpcreflect.NewStaticReflector(
			"acme.user.v1.UserService", // todo pass in full.Name for all services here
			"acme.group.v1.GroupService",
			// protoc-gen-connect-go generates package-level constants
			// for these fully-qualified protobuf service names, so you'd more likely
			// reference userv1.UserServiceName and groupv1.GroupServiceName.
		  )
	*/
	// mux.Handle(grpcreflect.NewHandlerV1(reflector))

	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(sampleconnect.NewExampleServiceHandler(&exampleservice.ExampleService{}))

	mux.Handle(sampleconnect.NewStreamingServiceHandler(&streamingservice.StreamingService{}))

	err := http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
