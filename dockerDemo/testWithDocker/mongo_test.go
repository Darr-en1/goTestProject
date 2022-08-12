package testWithDocker

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goTestProject/dockerDemo/commons"
	"log"
	"reflect"
	"sync"
	"testing"
)

var (
	once       sync.Once
	collection *mongo.Collection
)

func getAccountCollection(ctx context.Context) *mongo.Collection {
	once.Do(func() {
		connect, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Fatalf("cannot connect mongo: %+v", err)
		}
		collection = connect.Database("coolar").Collection("account")
	})
	return collection
}

// 表格驱动测试 golang 推荐的测试方式，输入输出非常直观

func TestMongo_InsertManyAccount(t *testing.T) {
	c := context.Background()
	collection := getAccountCollection(c)

	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		ctx       context.Context
		documents []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "InsertManyAccount",
			fields: fields{Collection: collection},
			args: args{
				ctx: c,
				documents: []interface{}{
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
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mongo{
				Collection: tt.fields.Collection,
			}
			if err := m.InsertManyAccount(tt.args.ctx, tt.args.documents); (err != nil) != tt.wantErr {
				t.Errorf("InsertManyAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongo_FindOneAccount(t *testing.T) {
	c := context.Background()
	collection := getAccountCollection(c)

	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "FindOneAccount darr_en1_001",
			fields: fields{Collection: collection},
			args: args{
				ctx: c,
				filter: bson.M{
					"open_id": "open_id_001",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mongo{
				Collection: tt.fields.Collection,
			}
			_, err := m.FindOneAccount(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOneAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMongo_ResolveAccountWithID(t *testing.T) {
	c := context.Background()
	collection := getAccountCollection(c)

	idObj := primitive.NewObjectID()
	newIdObj := primitive.NewObjectID()

	type fields struct {
		Collection *mongo.Collection
	}
	type args struct {
		ctx       context.Context
		openID    string
		documents interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *commons.Account
		wantErr bool
	}{
		{
			name:   "ResolveAccountWithID not exist",
			fields: fields{Collection: collection},
			args: args{
				ctx:    c,
				openID: "open_id_003",
				documents: bson.M{
					commons.IDField:         idObj,
					commons.NameField:       "darr_en1_003",
					commons.LoginCountField: 0,
				},
			},
			want: &commons.Account{
				ID:         idObj,
				OpenID:     "open_id_003",
				Name:       "darr_en1_003",
				LoginCount: 0,
			},
			wantErr: false,
		},
		{
			name:   "ResolveAccountWithID exist",
			fields: fields{Collection: collection},
			args: args{
				ctx:    c,
				openID: "open_id_003",
				documents: bson.M{
					commons.IDField:         newIdObj,
					commons.NameField:       "darr_en1_003_modify",
					commons.LoginCountField: 0,
				},
			},
			want: &commons.Account{
				ID:         idObj,
				OpenID:     "open_id_003",
				Name:       "darr_en1_003",
				LoginCount: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mongo{
				Collection: tt.fields.Collection,
			}
			got, err := m.ResolveAccountWithID(tt.args.ctx, tt.args.openID, tt.args.documents)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveAccountWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveAccountWithID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
