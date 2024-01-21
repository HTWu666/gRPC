package main

import (
	"context"
	"fmt"
	pb "grpc/server/grpcServer/proto"
	"log"
	"os"
	"flag"
	"net/http"
	"io"
	"strings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	flag.Parse()
	r := gin.Default()

	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER_IP") + ":" + os.Getenv("GRPC_SERVER_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTransmitClient(conn)

	r.GET("/grpc", func(c *gin.Context) {
		fileContent, err := os.ReadFile("./file.txt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		resp, err := client.Transmit(context.Background(), &pb.TransmitRequest{Request: string(fileContent)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send gRPC request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": resp.GetResponse()})
	})

	r.GET("/http", func(c *gin.Context) {
		fileContent, err := os.ReadFile("./file.txt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		resp, err := http.Post("http://" + os.Getenv("HTTP_SERVER_IP") + ":" + os.Getenv("HTTP_SERVER_PORT"), "text/plain", strings.NewReader(string(fileContent)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send HTTP request"})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("Request failed with status: %s", resp.Status)})
			return
		}

		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		c.String(http.StatusOK, string(responseData))
	})

	r.Run(":" + os.Getenv("CLIENT_PORT"))
}