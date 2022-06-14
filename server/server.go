package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/hanbarfe/grpc_example/gen/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Server struct {
	pb.UnimplementedExampleApiServer
}

func (t *Server) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received message: %v", in.GetName())
	userId := int32(rand.Intn(100))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: userId}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterExampleApiServer(grpcServer, &Server{})
	log.Printf("Server is listening at port: %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
