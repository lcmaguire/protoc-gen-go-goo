package exampleservice

import (
	example "protoc-gen-go-goo/out/example"
)

type ExampleService struct {
	example.UnimplementedExampleServiceServer
}
