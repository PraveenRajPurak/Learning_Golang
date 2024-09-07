package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Hello, World!")

	err := godotenv.Load()
	if err != nil {

		log.Fatalln(err)
	}

	URI := os.Getenv("MONGODB_URI")

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	client := connection(URI)

	if client != nil {

		fmt.Println("Connected to MongoDB! Truely")
	}

	router.Run()
}

func connection(URI string) *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))

	if err != nil {

		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {

		log.Fatalln(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client

}
