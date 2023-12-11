package exampleservice

import (
	"context"
	connect "connectrpc.com/connect"
	sample "github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// UpdateExample implements tutorial.ExampleService.UpdateExample.
func (s *Service) UpdateExample(ctx context.Context, req *connect.Request[sample.Example]) (*connect.Response[sample.Example], error) {
	_, err := s.firestore.Doc(req.Msg.Name).Set(ctx, req.Msg) // .Update may be useful with FieldMask.
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}
