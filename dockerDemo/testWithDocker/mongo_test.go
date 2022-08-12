package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/commons"
	"goTestProject/dockerDemo/testEnvironment"
	"os"
	"testing"
)

var mongoURI string

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	connect, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongo :v%", err)
	}
	m := NewMongo(connect.Database("coolar"))
	collection := m.Collection("account")
	id = commons.MustObjID("1231231221313131313")
	row, err := m.ResolveAccountID(c, collection, id)
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else {
		if row.ID != id {
			t.Errorf("get account id: want：%q; got：%q", want, row.ID.Hex())
		}
	}

}

// TestMain test  开始执行, 结束终止
func TestMain(m *testing.M) {
	os.Exit(testEnvironment.RunWithMongoInDocker(m, &mongoURI))
}
