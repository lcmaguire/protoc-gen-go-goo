package exampleservice

import (
	"github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	sampleconnect.UnimplementedExampleServiceHandler
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
