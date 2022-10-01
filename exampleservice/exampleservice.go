package exampleservice

import (
	out "github.com/lcmaguire/protoc-gen-go-goo/example/out"
)

// ExampleService ...
type ExampleService struct {
	out.UnimplementedExampleServiceServer
}

// test file:{name:"exampleservice/exampleservice.go"  content:"package exampleservice\n\nimport (\n\tout \"github.com/lcmaguire/protoc-gen-go-goo/example/out\"\n)\n\n// ExampleService ...\ntype ExampleService struct {\n\tout.UnimplementedExampleServiceServer\n}\n"}
