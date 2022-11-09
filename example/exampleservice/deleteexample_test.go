package exampleservice

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &example.SearchRequest{}
	res, err := service.DeleteExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res, &emptypb.Empty{})
}
