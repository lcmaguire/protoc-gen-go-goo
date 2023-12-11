package exampleservice

import (
	"context"
	"errors"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// GetExample implements tutorial.ExampleService.GetExample.
func (s *ExampleService) GetExample(ctx context.Context, req *connect.Request[sample.GetExampleRequest]) (*connect.Response[sample.Example], error) {
	res := connect.NewResponse(&sample.Example{})
	return res, connect.NewError(connect.CodeUnimplemented, errors.New("not yet implemented"))
}
