package exampleservice

import (
	context "context"
	example "protoc-gen-go-goo/out/example"
)

func (e *ExampleService) Example(ctx context.Context, in *example.SearchRequest) (out *example.SearchResponse, err error) {
	return
}
