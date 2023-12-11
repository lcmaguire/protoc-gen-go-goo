package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (s *Service) ListExamples(ctx context.Context, req *connect.Request[sample.ListExampleRequest]) (*connect.Response[sample.ListExampleResponse], error) {
	docSnaps, err := s.firestore.Collection("testCollection").Documents(ctx).GetAll() // hardcoding collection for now. Should probably be MessageName plural.
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	arr := make([]*sample.Example, 0)
	for _, v := range docSnaps {
		if v == nil || v.Data() == nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		var data *sample.Example
		if err := v.DataTo(&data); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		arr = append(arr, data)
	}
	return connect.NewResponse(
		&sample.ListExampleResponse{
			Examples: arr,
		},
	), nil
}
