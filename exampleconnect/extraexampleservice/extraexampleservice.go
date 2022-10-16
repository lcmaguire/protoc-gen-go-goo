package extraexampleservice

import (
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
)

// ExtraExampleService implements tutorial.ExtraExampleService.
type ExtraExampleService struct {
	example.UnimplementedExtraExampleServiceServer
}

func NewExtraExampleService() *ExtraExampleService {
	return &ExtraExampleService{}
}
