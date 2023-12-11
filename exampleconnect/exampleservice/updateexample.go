package exampleservice

import (
	"context"
	"errors"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (s *ExampleService) UpdateExample(ctx context.Context, req *connect.Request[sample.Example]) (*connect.Response[sample.Example], error) {
	res := connect.NewResponse(&sample.Example{})
	return res, connect.NewError(connect.CodeUnimplemented, errors.New("not yet implemented"))
}
