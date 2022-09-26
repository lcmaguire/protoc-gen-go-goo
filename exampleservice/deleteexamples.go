package exampleservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/out/example"
)

func (e *ExampleService) DeleteExamples(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return
}
