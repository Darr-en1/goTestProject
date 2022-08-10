package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id"`
	OpenID     string             `bson:"open_id"`
	Name       string             `bson:"name"`
	LoginCount int32              `bson:"login_count"`
}

func main() {
	c := context.Background()
	connect, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := connect.Database("coolcar").Collection("account")
	//insertRows(c, collection)
	findRow(c, collection)
}

func insertRows(ctx context.Context, collection *mongo.Collection) {
	res, err := collection.InsertMany(ctx, []interface{}{
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
	fmt.Printf("%+v", res)
}

func findOneRow(ctx context.Context, collection *mongo.Collection) {
	res := collection.FindOne(ctx, bson.M{
		"open_id": "1234",
	})
	var row Account
	err := res.Decode(&row)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", row)
}

func findRow(ctx context.Context, collection *mongo.Collection) {
	cur, err := collection.Find(ctx, bson.M{
		"open_id": "1234",
	})
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var row Account
		err := cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", row)
	}
}
