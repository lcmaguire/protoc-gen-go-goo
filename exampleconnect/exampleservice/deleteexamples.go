package exampleservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sample "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sample"
)

// DeleteExamples implements tutorial.ExampleService.DeleteExamples.
func (e *ExampleService) DeleteExamples(ctx context.Context, req *connect_go.Request[sample.SearchRequest]) (*connect_go.Response[sample.SearchResponse], error) {
	res := connect_go.NewResponse(&sample.SearchResponse{})
	return res, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("not yet implemented"))
}
