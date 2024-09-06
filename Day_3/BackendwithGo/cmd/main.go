package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/driver"
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/handlers"
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var app config.GoAppTools
var validate *validator.Validate

func main() {

	gob.Register(map[string]interface{}{})
	gob.Register(primitive.NewObjectID())

	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger

	validate = validator.New()

	app.Validate = validate

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

	GoApp := handlers.NewGoApp(&app, client)

	Routes(appRouter, GoApp)

	err = appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}

	app.InfoLogger.Println("Server started !")

}
