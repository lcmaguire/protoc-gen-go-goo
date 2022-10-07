package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExampleService ...
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}
