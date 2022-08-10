package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	connect, err := mongo.Connect(c, options.Client().ApplyURI("mongodb：//localhostL27"))
	if err != nil {
		panic(err)
	}
	collection := connect.Database("coolcar").Collection("account")
	res, err := collection.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "1234",
		},
		bson.M{
			"open_id": "343555",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
