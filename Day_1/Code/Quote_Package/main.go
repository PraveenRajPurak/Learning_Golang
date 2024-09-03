package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Glass())
}

// I learnt how to import an external package in go.
// 1. Simply import the package and use it in your code.
// 2. Use terminal to run "go mod tidy" to update go.mod file and add go.sum file. These specify the necessary packages for the code to work.
// 3. Run "go run main.go" to run code.
