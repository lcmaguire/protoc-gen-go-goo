package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// GetExample implements tutorial.ExampleService.GetExample.
func (s *Service) GetExample(ctx context.Context, req *connect.Request[sample.GetExampleRequest]) (*connect.Response[sample.Example], error) {
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	res := &sample.Example{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(res), nil
}
