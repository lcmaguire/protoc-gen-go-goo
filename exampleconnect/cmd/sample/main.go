package main

import (
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleservice"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
	streamingservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/streamingservice"
	http2 "golang.org/x/net/http2"
	h2c "golang.org/x/net/http2/h2c"
	log "log"
	http "net/http"
)

func main() {
	mux := http.NewServeMux()
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