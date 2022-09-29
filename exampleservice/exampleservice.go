package exampleservice

import (
	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
)

// ExampleService ...
type ExampleService struct {
	out.UnimplementedExampleServiceServer
}
