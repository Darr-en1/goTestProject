package main

import (
	"context"
	"fmt"
	trippb "goTestProject/grpc-demo/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}(conn)
	client := trippb.NewTripServiceClient(conn)
	trip, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "1111111"})
	if err != nil {
		log.Fatalf("GetTrip error: %v", err)
	}
	fmt.Println(trip)

}
