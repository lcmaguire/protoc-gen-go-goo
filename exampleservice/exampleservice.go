package exampleservice

import (
	out "example/out"
)

type ExampleService struct {
	out.UnimplementedExampleServiceServer
}
