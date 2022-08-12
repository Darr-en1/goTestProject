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
	OpenIDField     = "open_id"
	IDField         = "_id"
	NameField       = "name"
	LoginCountField = "login_count"
)

// Set $set 如果结果存在但是修改的项包含_id 是不被允许的
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert $setOnInsert  如果结果存在会修改除了_id 以外的其它行
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

func MustObjID(id string) primitive.ObjectID {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return hex
}
