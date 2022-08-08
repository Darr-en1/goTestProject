package main

import (
	trippb "goTestProject/grpc-demo/gen/go"
	trip "goTestProject/grpc-demo/tripService"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
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
