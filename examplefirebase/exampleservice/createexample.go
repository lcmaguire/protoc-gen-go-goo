package exampleservice

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	// need to set _id to resource name, or filter from resource, name
	mongoInsert, err := s.mongo.Database("local").Collection("colly").InsertOne(context.Background(), sample.Example{
		Name:        "nombre",
		DisplayName: "dippy",
	})
	if err != nil {
		fmt.Println("write err.")
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	/*
		bsonProto, err := bson.Marshal(&sample.Example{})
		if err != nil {
			return nil, err
		}
	*/
	//projection := &options.FindOneOptions{Projection: bsonProto}

	//opts := options.Find().SetProjection(bson.D{{"_id", ""}})

	type mongMessage struct {
		_id         string `bson:"_id" json:"id"`
		name        string
		DisplayName string
	}

	projection := options.FindOne().SetProjection(bson.D{{"_id", ""}})
	sing := s.mongo.Database("local").Collection("colly").FindOne(context.Background(), mongoInsert.InsertedID, projection)
	if sing == nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}
	if sing.Err() != nil {
		fmt.Println("singerr")
		return nil, sing.Err()
	}
	fmt.Println(sing)
	var protoRes *sample.Example
	// https://www.mongodb.com/docs/manual/reference/operator/projection/ specifies fields to return
	// https://www.mongodb.com/docs/manual/reference/operator/

	err = sing.Decode(protoRes)
	if err != nil {
		fmt.Println("decodyBytes")
		return nil, err
	}

	res := connect_go.NewResponse(protoRes) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}

// https://www.mongodb.com/docs/drivers/go/current/quick-start/
