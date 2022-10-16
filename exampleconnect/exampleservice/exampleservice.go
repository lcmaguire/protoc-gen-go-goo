package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
