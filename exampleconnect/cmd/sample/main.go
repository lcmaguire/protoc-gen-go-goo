// comment

package main

import (
	"log"
	"net/http"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"

	// your protoPathHere
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"

	// your services
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleservice"
)

func main() {
	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		"tutorial.ExampleService",
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(sampleconnect.NewExampleServiceHandler(&exampleservice.ExampleService{}))

	err := http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
