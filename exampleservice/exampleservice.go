package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/out/example"
)

type ExampleService struct {
	example.UnimplementedExampleServiceServer
}
