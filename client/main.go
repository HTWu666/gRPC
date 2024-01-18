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
)

func main() {
	flag.Parse()
	r := gin.Default()

	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

		// 创建请求
		resp, err := http.Post("http://localhost:5000", "text/plain", strings.NewReader(string(fileContent)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send HTTP request"})
			return
		}
		defer resp.Body.Close()

		// 检查响应状态
		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("Request failed with status: %s", resp.Status)})
			return
		}

		// 读取响应
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}
fmt.Println(string(responseData))
		// 将响应发送回客户端
		c.String(http.StatusOK, string(responseData))
	})

	r.Run(":3000")
}