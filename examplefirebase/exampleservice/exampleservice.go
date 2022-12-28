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
	// may need some path for collection + UID stuff passed in for user controlled data.
	Get(ctx context.Context, name string) (T, error)
	List(ctx context.Context) ([]T, error) // todo opts + for return list params
	Delete(ctx context.Context, name string) (T, error)
	Create(ctx context.Context, name string, msg T) (T, error)
	Update(ctx context.Context, name string, msg T) (T, error) // todo fieldmask
}

// gen below when message firebase included in cfg for plugin
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

	if err := docSnap.DataTo(&res); err != nil {
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return res, err
}

func (f *FirestoreDb[T]) List(ctx context.Context) (res []T, err error) {
	// collection path should be const
	docSnaps, err := f.firestore.Collection("testCollection").Documents(ctx).GetAll() // hardcoding collection for now. Should probably be MessageName plural.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, err)
		}

		var data T
		if err := v.DataTo(&data); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, err)
		}
		res = append(res, data)
	}
	return res, nil
}

func (f *FirestoreDb[T]) Create(ctx context.Context, name string, in T) (res T, err error) {
	_, err = f.firestore.Doc(name).Create(ctx, in)
	if err != nil {
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return res, nil
}

func (f *FirestoreDb[T]) Delete(ctx context.Context, name string) (res T, err error) {
	_, err = f.firestore.Doc(name).Delete(ctx)
	if err != nil {
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return res, nil
}

func (f *FirestoreDb[T]) Update(ctx context.Context, name string, in T) (res T, err error) {
	_, err = f.firestore.Doc(name).Set(ctx, in)
	if err != nil {
		return res, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return in, nil
}
