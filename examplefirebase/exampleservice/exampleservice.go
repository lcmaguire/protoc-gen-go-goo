package exampleservice

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	auth "firebase.google.com/go/v4/auth"
	"github.com/golang/protobuf/proto"
	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
)

// Service implements tutorial.ExampleService.
type Service struct {
	sampleconnect.UnimplementedExampleServiceHandler
	firestore *firestore.Client
	auth      *auth.Client
}

func NewService(auth *auth.Client, firestore *firestore.Client) *Service {
	return &Service{
		auth:      auth,
		firestore: firestore,
	}
}

type Database interface {
	Get(ctx context.Context, name string) (proto.Message, error)
	List(ctx context.Context) (proto.Message, error) // todo opts
	Delete(ctx context.Context, name string) (proto.Message, error)
	Create(ctx context.Context, msg proto.Message) (proto.Message, error)
	Update(ctx context.Context, msg proto.Message) (proto.Message, error) // todo fieldmask
}

type FirestoreDb struct {
	Database
	firestore *firestore.Client
}

func (f *FirestoreDb) Get(ctx context.Context, name string) sample.Example {
	//
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}

	res := &sample.Example{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return res, err
}
