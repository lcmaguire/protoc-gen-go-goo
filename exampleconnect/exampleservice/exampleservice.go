package exampleservice

import (
	exampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
