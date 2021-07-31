package parser

import (
	"goTestProject/crewler/engine"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	type args struct {
		contents []byte
	}
	tests := []struct {
		name string
		args args
		want engine.ParseResult
	}{
		// TODO: Add test cases.
		{
			name: "zhenai",
			args: args{
				contents: body,
			},
			want: engine.ParseResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCityList(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCityList() = %v, want %v", got, tt.want)
			}
		})
	}
}
