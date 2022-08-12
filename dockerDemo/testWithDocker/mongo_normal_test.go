package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/commons"
	"testing"
)

//var mongoURI string

// 一般的test 编写方式，不清晰，不推荐

// IgnoreTestResolveAccountID 想要允许将名称改成TestResolveAccountID
func IgnoreTestResolveAccountID(t *testing.T) {
	c := context.Background()
	connect, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongo: %+v", err)
	}
	m := NewMongo(connect.Database("coolar").Collection("account"))

	err = m.InsertManyAccount(c, []interface{}{
		// 可以直接使用对象，需要声明 序列化映射  `bson:"_id"`,需要手动添加 _id 否则 会使用默认值 000000000000， 存在则报错
		commons.Account{
			ID:         primitive.NewObjectID(),
			OpenID:     "open_id_001",
			Name:       "darr_en1_001",
			LoginCount: 0,
		},
		bson.M{
			commons.OpenIDField:     "open_id_002",
			commons.NameField:       "darr_en1_002",
			commons.LoginCountField: 0,
		},
	})
	if err != nil {
		t.Errorf("InsertManyAccount error: %v", err)
	}

	account, err := m.FindOneAccount(c, bson.M{
		"open_id": "open_id_001",
	})
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else if account.Name != "darr_en1_001" {
		t.Errorf("get account name: want：%q; got：%q", "name_001", account.Name)
	}

	openID := "open_id_003"
	row, err := m.ResolveAccount(c, openID, bson.M{commons.NameField: "darr_en1_003"})
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else if row.Name != "darr_en1_003" {
		// %q 输出带引号
		t.Errorf("get account name: want：%q; got：%q", "darr_en1_003", row.Name)
	}

	idObj := primitive.NewObjectID()
	openID = "open_id_004"
	name := "darr_en1_004"
	row, err = m.ResolveAccountWithID(c, openID, bson.M{
		commons.IDField:         idObj,
		commons.OpenIDField:     openID,
		commons.NameField:       name,
		commons.LoginCountField: 0,
	})
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else if row.ID != idObj {
		// %q 输出带引号
		t.Errorf("get account id: want：%q; got：%q", idObj.Hex(), row.ID.Hex())
	}

	// open_id 对应的_id 不一致 不更新
	newIdObj := primitive.NewObjectID()
	row, err = m.ResolveAccountWithID(c, openID, bson.M{
		commons.IDField:         newIdObj,
		commons.OpenIDField:     openID,
		commons.NameField:       name,
		commons.LoginCountField: 0,
	})
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else if row.ID != idObj {
		// %q 输出带引号
		t.Errorf("get account id: want：%q; got：%q", idObj.Hex(), row.ID.Hex())
	}

}
