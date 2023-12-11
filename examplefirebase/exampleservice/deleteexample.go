package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DeleteExample implements tutorial.ExampleService.DeleteExample.
func (s *Service) DeleteExample(ctx context.Context, req *connect.Request[sample.DeleteExampleRequest]) (*connect.Response[emptypb.Empty], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Delete(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}
