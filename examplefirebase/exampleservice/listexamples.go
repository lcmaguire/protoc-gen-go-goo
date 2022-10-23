package exampleservice

import (
	context "context"
	errors "errors"
	"fmt"

	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (e *ExampleService) ListExamples(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) { // would need to change to
	docSnaps, err := e.firestore.Collection("testCollection").Documents(ctx).GetAll() // todo get uid from request.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
	}

	res := &sample.SearchResponse{}
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeNotFound, errors.New("err not found"))
		}
		// DataAt(path string) or DataAtPath(fp FieldPath) look fun
		if err := v.DataTo(res); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("asdsadf"))
		}
		fmt.Println(res)
	}

	return connect_go.NewResponse(res), nil
}
