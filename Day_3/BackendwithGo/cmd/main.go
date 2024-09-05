package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/driver"
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app config.GoAppTools

func main() {

	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger

	fmt.Println("Creating a server !")

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}
	app.InfoLogger.Println("Connecting to MongoDB!")

	URI := os.Getenv("MONGODB_URI")

	if URI == "" {
		app.ErrorLogger.Fatalln("MONGO_URI not found in .env file")
	}

	client := driver.Connection(URI, app)

	defer func() {

		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatalln(err)
			return
		}
	}()

	appRouter := gin.New()

	appRouter.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Welcome to the Backend with Go",
		})
	})

	err = appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}

	app.InfoLogger.Println("Server started !")

}
