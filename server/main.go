package main

import (
	"context"
	"fmt"
	pb "grpc/server/proto"
	"net"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTransmitServer
}

func (s *server) Transmit(ctx context.Context, req *pb.TransmitRequest) (*pb.TransmitResponse, error) {
	fmt.Println("grpc is running")
	return &pb.TransmitResponse{Response: "Server received your request"}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":3000")
	grpcServer := grpc.NewServer()
	pb.RegisterTransmitServer(grpcServer, &server{})
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("Failed to serve: %v", err)
		return
	}
}