package main

import (
	exampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleconnect"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/exampleservice"
	extraexampleservice "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/extraexampleservice"
	http2 "golang.org/x/net/http2"
	h2c "golang.org/x/net/http2/h2c"
	log "log"
	http "net/http"
)

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.

	mux.Handle(exampleconnect.NewExampleServiceHandler(&exampleservice.ExampleService{}))

	mux.Handle(exampleconnect.NewExtraExampleServiceHandler(&extraexampleservice.ExtraExampleService{}))

	err := http.ListenAndServe(
		"localhost:8080",
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
