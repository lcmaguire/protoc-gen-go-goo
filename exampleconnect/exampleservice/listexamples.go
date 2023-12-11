package exampleservice

import (
	"context"
	"errors"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (s *ExampleService) ListExamples(ctx context.Context, req *connect.Request[sample.ListExampleRequest]) (*connect.Response[sample.ListExampleResponse], error) {
	res := connect.NewResponse(&sample.ListExampleResponse{})
	return res, connect.NewError(connect.CodeUnimplemented, errors.New("not yet implemented"))
}
