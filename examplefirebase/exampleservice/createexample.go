package exampleservice

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (s *Service) CreateExample(ctx context.Context, req *connect_go.Request[sample.Example]) (*connect_go.Response[sample.Example], error) {
	_, err := s.db.Create(ctx, req.Msg.Name, req.Msg)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
