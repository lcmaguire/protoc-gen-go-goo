package exampleservice

import (
	"context"
	"errors"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (s *ExampleService) DeleteExample(ctx context.Context, req *connect.Request[sample.DeleteExampleRequest]) (*connect.Response[emptypb.Empty], error) {
	res := connect.NewResponse(&emptypb.Empty{})
	return res, connect.NewError(connect.CodeUnimplemented, errors.New("not yet implemented"))
}
