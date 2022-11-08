package exampleservice

import (
	"context"
	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (s *Service) UpdateExample(ctx context.Context, req *connect_go.Request[sample.Example]) (*connect_go.Response[sample.Example], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
