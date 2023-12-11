package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

func TestGetExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect.Request[sample.GetExampleRequest]{
		Msg: &sample.GetExampleRequest{},
	}
	res, err := service.GetExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect.CodeUnimplemented, connect.CodeOf(err))
	proto.Equal(res.Msg, &sample.Example{})
}
