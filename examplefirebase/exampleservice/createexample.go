package exampleservice

import (
	context "context"
	errors "errors"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (e *ExampleService) CreateExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	_, err := e.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
	if err != nil {
		connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}
	res := connect_go.NewResponse(&sample.SearchResponse{Name: docRef.Path})
	return res, nil
}
