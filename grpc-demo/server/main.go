package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	trippb "goTestProject/grpc-demo/gen/go"
	trip "goTestProject/grpc-demo/tripService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	go startGRPCGateway()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGRPCGateway() {
	c := context.Background()
	// 创建可以cancel 的 上下文
	c, cancel := context.WithCancel(c)
	defer cancel()

	// 分发器，将http请求分发到grpc server
	mux := runtime.NewServeMux()

	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("cannot grpc gateway: %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot Listen and server: %v", err)
	}
}
