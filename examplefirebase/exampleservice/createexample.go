package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (s *Service) CreateExample(ctx context.Context, req *connect.Request[sample.Example]) (*connect.Response[sample.Example], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
