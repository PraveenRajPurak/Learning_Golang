package main

import (
	"fmt"
	"log"

	"hello.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Greetings("Praveen")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
