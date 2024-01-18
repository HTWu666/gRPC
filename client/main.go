package main

import (
	"context"
	"fmt"
	pb "grpc/server/proto"
	"log"
	"os"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransmitClient(conn)

	fileContent, err := os.ReadFile("./file.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	resp, err := client.Transmit(context.Background(), &pb.TransmitRequest{Request: string(fileContent)})
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	fmt.Println(resp.GetResponse())
}