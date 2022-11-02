package exampleservice

import (
	"context"
	"errors"
	connect_go "github.com/bufbuild/connect-go"
	// sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"

	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (e *ExampleService) DeleteExample(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[emptypb.Empty], error) {
	res := connect_go.NewResponse(&emptypb.Empty{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
