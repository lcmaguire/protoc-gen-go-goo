package exampleservice

import (
	context "context"
	example "github.com/lcmaguire/protoc-gen-go-goo/out/example"
)

func (e *ExampleService) CreateExample(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return
}
