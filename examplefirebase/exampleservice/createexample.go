package exampleservice

import (
	"context"
	"fmt"
	"os"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// CreateExample implements tutorial.ExampleService.CreateExample.
func (s *Service) CreateExample(ctx context.Context, req *connect_go.Request[sample.Example]) (*connect_go.Response[sample.Example], error) {

	_, err := s.firestore.Doc(req.Msg.Name).Create(ctx, req.Msg)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	mongo.NewClient()
	res := connect_go.NewResponse(req.Msg) // hard coding for now assuming req and res type are same and Write is always successful.
	return res, nil
}

// https://www.mongodb.com/docs/drivers/go/current/quick-start/
func mongoManiac() (*mongo.Client, error) {

	// mongodb://root:example@mongo:27017/

	return mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@mongo:27017/"))
}

func mongito() {
	// uri := "mongodb://root:example@mongo:27017/"

	/*cli, err := mongoManiac()
	if err != nil {
		panic(err)
	}
	*/
	fmt.Println("aqui")
	fmt.Println(os.Getenv("MONGODB_URI"))

	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@127.0.0.1"))
	if err != nil {
		panic(err)
	}

	err = cli.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	/*
		db := cli.Database("local")
		err = db.CreateCollection(context.Background(), "colly")
		if err != nil {
			panic(err)
		}

		res, err := db.Collection("colly").InsertOne(context.Background(), sample.Example{
			Name:        "nombre",
			DisplayName: "dippy",
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(res)

		sing := db.Collection("colly").FindOne(context.Background(), res.InsertedID)
		fmt.Println(sing)
		// cli.ListDatabases()
	*/
}
