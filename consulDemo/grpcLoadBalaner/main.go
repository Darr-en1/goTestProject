package main

import (
	"context"
	"fmt"
	trippb "goTestProject/grpc-demo/gen/go"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	// 使用grpc 通过 consul 的负载均衡 获取 server conn
	conn, err := grpc.Dial(
		"consul://172.25.40.121:8500/grpcServer?wait=14s",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // pick_first（默认）round_robin、 和 grpclb
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := trippb.NewTripServiceClient(conn)
	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	trip, err := client.GetTrip(ctx, &trippb.GetTripRequest{Id: "1111111"})
	if err != nil {
		log.Fatalf("GetTrip error: %v", err)
	}
	fmt.Println(trip)
}
