package exampleservice

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/lcmaguire/protoc-gen-go-goo/example"
)

func TestCreateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &example.SearchRequest{}
	res, err := service.CreateExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res, &example.SearchResponse{})
}
