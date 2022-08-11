package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id"`
	OpenID     string             `bson:"open_id"`
	Name       string             `bson:"name"`
	LoginCount int32              `bson:"login_count"`
}

func FindOneRow(ctx context.Context, collection *mongo.Collection) (*Account, error) {
	res := collection.FindOne(ctx, bson.M{
		"open_id": "1234",
	})
	var acc Account
	err := res.Decode(&acc)
	return &acc, err
}
