package exampleservice

import (
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sampleconnect"
)

// ExampleService implements tutorial.ExampleService.
type ExampleService struct {
	sampleconnect.UnimplementedExampleServiceHandler
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}
