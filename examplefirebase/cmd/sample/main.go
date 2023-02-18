package main

import (
	"context"
	v4 "firebase.google.com/go/v4"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/api/option"
	"log"
	"net/http"
	// your protoPathHere
	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
	// your services
	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/exampleservice"
)

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(sampleconnect.NewExampleServiceHandler(createNewService()))
	err := http.ListenAndServe(
		"localhost:8080", // auth host users 8080
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}

// createNewService creates a new Service, exampleservice pkg is hard coded for now
func createNewService() *exampleservice.Service {
	opt := option.WithCredentialsFile("your-firebase-service-account.json") // todo have this be env var
	app, err := v4.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	firestore, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return exampleservice.NewService(auth, firestore)
}
