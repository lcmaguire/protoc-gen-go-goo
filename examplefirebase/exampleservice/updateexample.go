package exampleservice

import (
	context "context"
	errors "errors"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (e *ExampleService) UpdateExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {

	/*
		req.Msg.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name("")) + req.fieldMas[x] to use below
		_, err := e.firestore.Doc(req.Msg.Name).Update(ctx, []firestore.Update{})
	*/
	// set will act as complete replacement.
	_, err := e.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	docSnap, err := e.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, errors.New("err not found"))
	}

	res := &sample.SearchResponse{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("unmarshal"))
	}

	return connect_go.NewResponse(res), nil
}
