package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/testEnvironment"
	"os"
	"testing"
)

var mongoURI string

func TestFindOneRow(t *testing.T) {
	c := context.Background()
	connect, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	collection := connect.Database("coolcar").Collection("account")
	row, err := FindOneRow(c, collection)
	if err != nil {
		t.Errorf("faild get account error: %v", err)
	} else {
		want := "asasasasa"
		if row.ID.Hex() != want {
			t.Errorf("get account id: want：%q; got：%q", want, row.ID.Hex())
		}
	}

}

// TestMain test test 开始执行, 结束终止
func TestMain(m *testing.M) {
	os.Exit(testEnvironment.RunWithMongoInDocker(m, &mongoURI))
}
