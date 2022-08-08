package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	trippb "goTestProject/grpc-demo/gen/go"
)

// protoc -I=. --go_out=paths=source_relative:gen/go trip.proto 生成 go 代码，可以直接使用生成的 go struct
func main() {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 3600,
		FeeCent:     10000,
		StartPos: &trippb.Location{
			Latitude:   200,
			Longtitude: 300,
		},
		Status: trippb.TripStatus_NOT_STARTED,
	}
	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	// 16进制 输出
	fmt.Printf("%X\n", b)

	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	jsonB, err := json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonB)
}
