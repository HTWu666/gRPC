package main

import (
	"context"
	"fmt"
	pb "grpc/server/grpcServer/proto"
	"net"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTransmitServer
}

func (s *server) Transmit(ctx context.Context, req *pb.TransmitRequest) (*pb.TransmitResponse, error) {
	return &pb.TransmitResponse{Response: "Server received your request"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":5001")
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