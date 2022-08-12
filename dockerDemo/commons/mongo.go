package commons

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id"`
	OpenID     string             `bson:"open_id"`
	Name       string             `bson:"name"`
	LoginCount int32              `bson:"login_count"`
}

const (
	OpenIDField = "open_id"
	IDField     = "_id"
)

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert() bson.M {
	return
}

func MustObjID(id string) primitive.ObjectID {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return hex
}
