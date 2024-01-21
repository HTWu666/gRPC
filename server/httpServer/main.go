package main

import (
	"net/http"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server received your request",
		})
	})

	r.Run(":" + os.Getenv("HTTP_SERVER_PORT"))
}
