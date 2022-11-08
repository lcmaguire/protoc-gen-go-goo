package exampleservice

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (s *Service) DeleteExample(ctx context.Context, req *connect_go.Request[sample.DeleteExampleRequest]) (*connect_go.Response[emptypb.Empty], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Delete(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	return connect_go.NewResponse(&emptypb.Empty{}), nil
}
