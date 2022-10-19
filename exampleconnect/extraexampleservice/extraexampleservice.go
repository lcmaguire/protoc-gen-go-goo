package extraexampleservice

import (
	sampleconnect "github.com/lcmaguire/protoc-gen-go-goo/exampleconnect/sampleconnect"
)

// ExtraExampleService implements tutorial.ExtraExampleService.
type ExtraExampleService struct {
	sampleconnect.UnimplementedExtraExampleServiceHandler
}

func NewExtraExampleService() *ExtraExampleService {
	return &ExtraExampleService{}
}
