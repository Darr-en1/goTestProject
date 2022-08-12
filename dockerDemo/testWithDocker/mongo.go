package testWithDocker

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/commons"
)

type Mongo struct {
	*mongo.Collection
}

func NewMongo(collection *mongo.Collection) *Mongo {
	return &Mongo{collection}
}

func (m *Mongo) InsertManyAccount(ctx context.Context, documents []interface{}) error {
	res, err := m.InsertMany(ctx, documents)
	// 返回的是 _id 集合
	fmt.Println(res.InsertedIDs[0].(primitive.ObjectID).Hex())
	return err
}

func (m *Mongo) FindOneAccount(ctx context.Context, filter interface{}) (*commons.Account, error) {
	row := m.FindOne(ctx, filter)
	var res commons.Account
	err := row.Decode(&res)
	return &res, err
}

// ResolveAccount 不能修改 _id  因为它是不可变字段 如果添加了 _id，数据已存在且与传入的_id不对应则会报错
func (m *Mongo) ResolveAccount(ctx context.Context, openID string, documents interface{}) (*commons.Account, error) {

	res := m.FindOneAndUpdate(ctx,
		bson.M{
			commons.OpenIDField: openID,
		},
		commons.Set(documents),
		// SetUpsert true 没有则创建
		// SetReturnDocument After 返回修改后的内容
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	var acc commons.Account
	err := res.Decode(&acc)
	return &acc, err
}

// ResolveAccountWithID 如果传入的_id和查询的_id 一致 则修改，否则返回 open_id 查询的结果
func (m *Mongo) ResolveAccountWithID(ctx context.Context, openID string, documents interface{}) (*commons.Account, error) {

	res := m.FindOneAndUpdate(ctx,
		bson.M{
			commons.OpenIDField: openID,
		},
		commons.SetOnInsert(documents),
		// SetUpsert true 没有则创建
		// SetReturnDocument After 返回修改后的内容
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	var acc commons.Account
	err := res.Decode(&acc)
	return &acc, err
}
