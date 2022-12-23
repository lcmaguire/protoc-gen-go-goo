package exampleservice

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	auth "firebase.google.com/go/v4/auth"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/proto"
	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
)

// Service implements tutorial.ExampleService.
type Service struct {
	sampleconnect.UnimplementedExampleServiceHandler
	firestore *firestore.Client
	auth      *auth.Client
	db        Database[*sample.Example]
}

func NewService(auth *auth.Client, firestore *firestore.Client) *Service {
	return &Service{
		auth:      auth,
		firestore: firestore,
		db:        &FirestoreDb[*sample.Example]{firestore: firestore},
	}
}

type Database[T proto.Message] interface {
	Get(ctx context.Context, name string) (T, error)
	List(ctx context.Context) ([]T, error) // todo opts
	Delete(ctx context.Context, name string) (T, error)
	Create(ctx context.Context, msg T) (T, error)
	Update(ctx context.Context, msg T) (T, error) // todo fieldmask
}

type FirestoreDb[T proto.Message] struct {
	Database[T]
	firestore *firestore.Client
}

func (f *FirestoreDb[T]) Get(ctx context.Context, name string) (res T, err error) {
	docSnap, err := f.firestore.Doc(name).Get(ctx)
	if err != nil {
		fmt.Println("aqui")
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}

	if docSnap == nil || docSnap.Data() == nil {
		fmt.Println("no doc")
		return res, connect_go.NewError(connect_go.CodeNotFound, err)
	}
	var out T
	if err := docSnap.DataTo(&out); err != nil {
		fmt.Println("unable to parse")
		fmt.Println(out)
		fmt.Println(docSnap)
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return out, err
}
