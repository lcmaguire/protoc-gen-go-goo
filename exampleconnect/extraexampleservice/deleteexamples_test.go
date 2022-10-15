package extraexampleservice

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/golang/protobuf/proto"
	example "github.com/lcmaguire/protoc-gen-go-goo/example"
	assert "github.com/stretchr/testify/assert"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	testing "testing"
)

func TestDeleteExamples(t *testing.T) {
	t.Parallel()
	service := &ExtraExampleService{}
	req := &connect_go.Request[example.SearchRequest]{
		Msg: &example.SearchRequest{},
	}
	res, err := service.DeleteExamples(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, codes.Unimplemented, status.Code(err))
	proto.Equal(res.Msg, &emptypb.Empty{})
}
