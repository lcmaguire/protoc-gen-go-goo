package exampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExampleService ...
type ExampleService struct {
	example.UnimplementedExampleServiceServer
}

// test file:{name:"exampleservice/exampleservice.go" content:"package exampleservice\n\nimport (\n\texample \"github.com/lcmaguire/protoc-gen-go-goo/example\"\n)\n\n// ExampleService ...\ntype ExampleService struct {\n\texample.UnimplementedExampleServiceServer\n}\n"}
