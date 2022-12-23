package exampleservice

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// ListExamples implements tutorial.ExampleService.ListExamples.
func (s *Service) ListExamples(ctx context.Context, req *connect_go.Request[sample.ListExampleRequest]) (*connect_go.Response[sample.ListExampleResponse], error) {
	arr, err := s.db.List(ctx)
	if err != nil {
		return nil, err // todo wrap err
	}
	return connect_go.NewResponse(
		&sample.ListExampleResponse{
			Examples: arr,
		},
	), nil
}
