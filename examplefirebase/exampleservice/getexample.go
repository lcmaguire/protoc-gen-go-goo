package exampleservice

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// GetExample implements tutorial.ExampleService.GetExample.
func (s *Service) GetExample(ctx context.Context, req *connect_go.Request[sample.GetExampleRequest]) (*connect_go.Response[sample.Example], error) {
	docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, errors.New("err not found"))
	}

	res := &sample.Example{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, errors.New("error loading response"))
	}
	return connect_go.NewResponse(res), nil
}
