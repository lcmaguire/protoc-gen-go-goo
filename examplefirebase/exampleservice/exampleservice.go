package exampleservice

import (
	firestore "cloud.google.com/go/firestore"
	auth "firebase.google.com/go/v4/auth"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service implements tutorial.ExampleService.
type Service struct {
	sampleconnect.UnimplementedExampleServiceHandler
	firestore *firestore.Client
	mongo     *mongo.Client
	auth      *auth.Client
}

func NewService(auth *auth.Client, firestore *firestore.Client, mongo *mongo.Client) *Service {
	return &Service{
		auth:      auth,
		firestore: firestore,
		mongo:     mongo,
	}
}
