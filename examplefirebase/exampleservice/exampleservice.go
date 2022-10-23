package exampleservice

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4/auth"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	sampleconnect.UnimplementedExampleServiceHandler
	firestore *firestore.Client
	auth      *auth.Client
}

func NewExampleService(auth *auth.Client, firestore *firestore.Client) *ExampleService {
	return &ExampleService{
		auth:      auth,
		firestore: firestore,
	}
}
