package exampleservice

import (
	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
)

type ExampleService struct {
	out.UnimplementedExampleServiceServer
}
