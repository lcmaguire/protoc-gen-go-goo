package exampleservice

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

func TestUpdateExample(t *testing.T) {
	t.Parallel()
	service := &ExampleService{}
	req := &connect_go.Request[sample.Example]{
		Msg: &sample.Example{},
	}
	res, err := service.UpdateExample(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, connect_go.CodeUnimplemented, connect_go.CodeOf(err))
	proto.Equal(res.Msg, &sample.Example{})
}
