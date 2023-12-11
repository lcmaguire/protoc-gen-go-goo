package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect.Request[sample.DeleteExampleRequest]{
		Msg: &sample.DeleteExampleRequest{},
	}
	res, err := service.DeleteExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect.CodeUnimplemented, connect.CodeOf(err))
	proto.Equal(res.Msg, &emptypb.Empty{})
}
