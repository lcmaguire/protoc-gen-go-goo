package main

import (
	context "context"
	v4 "firebase.google.com/go/v4"
	exampleservice "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/exampleservice"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
	http2 "golang.org/x/net/http2"
	h2c "golang.org/x/net/http2/h2c"
	option "google.golang.org/api/option"
	log "log"
	http "net/http"
)

func main() {
	mux := http.NewServeMux()
	// The generated constructors return a path and a plain net/http
	// handler.
	mux.Handle(sampleconnect.NewExampleServiceHandler(createNewService()))
	// export FIREBASE_AUTH_EMULATOR_HOST="localhost:9099"
	// export FIRESTORE_EMULATOR_HOST="localhost:8080"
	err := http.ListenAndServe(
		"localhost:8080", // auth host users 8080
		// For gRPC clients, it's convenient to support HTTP/2 without TLS. You can
		// avoid x/net/http2 by using http.ListenAndServeTLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: " + err.Error())
}
func createNewService() *exampleservice.Service {
	opt := option.WithCredentialsFile("./test-firebase-service-account.json")
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
