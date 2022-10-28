package exampleservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (s *Service) ListExamples(ctx context.Context, req *connect_go.Request[sample.ListExampleRequest]) (*connect_go.Response[sample.ListExampleResponse], error) {
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // todo get uid from request.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}
	arr := []*sample.Example{}
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
		}

		var data *sample.Example
		if err := v.DataTo(&data); err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error unable to load response"))
		}
		arr = append(arr, data)
	}
	return connect_go.NewResponse(
		&sample.ListExampleResponse{
			Examples: arr,
		},
	), nil
}
