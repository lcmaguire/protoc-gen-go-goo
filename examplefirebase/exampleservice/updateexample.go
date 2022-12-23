package exampleservice

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (s *Service) UpdateExample(ctx context.Context, req *connect_go.Request[sample.Example]) (*connect_go.Response[sample.Example], error) {
	dbRes, err := s.db.Update(ctx, req.Msg.Name, req.Msg)
	if err != nil {
		return nil, err
	}

	res := connect_go.NewResponse(dbRes) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
