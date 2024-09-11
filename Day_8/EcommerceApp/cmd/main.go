package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/driver"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/handler"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/config"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
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

	fmt.Println("Welcome to Ecommerce App!")

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}
	URI := os.Getenv("MONGODB_URI")
	fmt.Println("MongoDB URI : ", URI)

	client := driver.Connection(URI, app)

	if client != nil {

		fmt.Println("Connected to MongoDB!")
	}

	webserver := gin.New()

	GoApp := handler.NewGoApp(&app, client)

	Routes(webserver, GoApp)

	webserver.Run(":10010")
}
