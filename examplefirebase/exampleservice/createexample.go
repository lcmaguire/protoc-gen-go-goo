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

func mongito() {
	// uri := "mongodb://root:example@mongo:27017/"

	/*cli, err := mongoManiac()
	if err != nil {
		panic(err)
	}
	*/
	fmt.Println("aqui")
	fmt.Println(os.Getenv("MONGODB_URI"))
	//uri := "mongodb://root:example@mongo:27017"
	//uri := "mongodb://root:example@mongo"
	uri := "mongodb://localhost:27017"
	fmt.Println(uri)
	//mongodb: //localhost:27017
	//uri := "mongodb://root:example@mongo:127.0.0.1:27017"
	/*
		try old without / at end
		try just mongo suffix
	*/
	// mongodb://root:example@127.0.0.1
	cli, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = cli.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	db := cli.Database("local")
	err = db.CreateCollection(context.Background(), "colly")
	if err != nil {
		fmt.Println(err)
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

}
