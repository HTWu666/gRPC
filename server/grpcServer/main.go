package main

import (
	"context"
	"fmt"
	pb "grpc/server/grpcServer/proto"
	"net"
	"google.golang.org/grpc"
	"github.com/joho/godotenv"
)

type server struct {
	pb.UnimplementedTransmitServer
}

func (s *server) Transmit(ctx context.Context, req *pb.TransmitRequest) (*pb.TransmitResponse, error) {
	return &pb.TransmitResponse{Response: "Server received your request"}, nil
}

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	listen, err := net.Listen("tcp", ":" + os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		fmt.Printf("Failed to connect: %v", err)
		return
	}

	fmt.Println("grpc is opening")

	grpcServer := grpc.NewServer()
	pb.RegisterTransmitServer(grpcServer, &server{})
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("Failed to serve: %v", err)
		return
	}
}