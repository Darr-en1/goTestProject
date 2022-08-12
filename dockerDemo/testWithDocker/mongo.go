package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/commons"
)

type Mongo struct {
	*mongo.Database
}

func NewMongo(dataBase *mongo.Database) *Mongo {
	return &Mongo{dataBase}
}

func (m *Mongo) ResolveAccountID(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*commons.Account, error) {
	var (
		openID = "1234"
	)

	res := collection.FindOneAndUpdate(ctx,
		bson.M{
			commons.OpenIDField: "1234",
		}, commons.SetOnInsert(bson.M{
			commons.IDField:     id,
			commons.OpenIDField: openID,
		}),
		// SetUpsert true 没有则创建
		// SetReturnDocument After 返回修改后的内容
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	var acc commons.Account
	err := res.Decode(&acc)
	return &acc, err
}
