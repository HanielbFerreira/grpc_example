package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hanbarfe/grpc_example/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = ":50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Something happened: %v", err)
	}

	defer conn.Close()
	client := pb.NewExampleApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateNewUser(ctx, &pb.NewUser{Name: "Haniel", Age: 33})
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	log.Printf(`User Details:
	NAME: %s
	AGE: %d
	ID: %d`, res.GetName(), res.GetAge(), res.GetId())

}
