package exampleservice

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (s *Service) CreateExample(ctx context.Context, req *connect_go.Request[sample.Example]) (*connect_go.Response[sample.Example], error) {
	/*
		_, err := s.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
		if err != nil {
			return nil, connect_go.NewError(connect_go.CodeInternal, err)
		}
	*/

	// create collection if it doesnt exist.

	// todo handle database creation elsewhere.
	mongoInsert, err := s.mongo.Database("local").Collection("colly").InsertOne(context.Background(), sample.Example{
		Name:        "nombre",
		DisplayName: "dippy",
	})
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	sing := s.mongo.Database("local").Collection("colly").FindOne(context.Background(), mongoInsert.InsertedID)
	if sing != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	protoRes := &sample.Example{}
	if err := sing.Decode(protoRes); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	res := connect_go.NewResponse(protoRes) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}

// https://www.mongodb.com/docs/drivers/go/current/quick-start/
