package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func main() {

	fmt.Println("Hello World!")

	webserver := gin.Default()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	URI := os.Getenv("MONGODB_URI")

	webserver.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	client, err := connection(URI)

	if err != nil {
		fmt.Println("Could not connect to mongodb!")
	}

	if client != nil {
		fmt.Println("Connected to MongoDB!")
		Client = client

		fmt.Println("Trying to insert data...")
	}

	webserver.POST("/user", InsertUser)

	webserver.GET("/users", GetUsers)

	webserver.Run(":10005")
}

func connection(URI string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func InsertUser(c *gin.Context) {

	var user User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Client.Database("testdb").Collection("user")

	user.ID = primitive.NewObjectID()

	insertionRes, err := collection.InsertOne(ctx, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": insertionRes})
}

func GetUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := Client.Database("testdb").Collection("user")

	filter := bson.D{{Key: "email", Value: "prp@gmail.com"}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
