package exampleservice

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// GetExample implements tutorial.ExampleService.GetExample.
func (s *Service) GetExample(ctx context.Context, req *connect_go.Request[sample.GetExampleRequest]) (*connect_go.Response[sample.Example], error) {
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}
	res := &sample.Example{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	return connect_go.NewResponse(res), nil
}
