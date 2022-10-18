package extraexampleservice

import (
	exampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect"
)

// ExtraExampleService implements tutorial.ExtraExampleService.
type ExtraExampleService struct {
	example.UnimplementedExtraExampleServiceServer
}

func NewExtraExampleService() *ExtraExampleService {
	return &ExtraExampleService{}
}
