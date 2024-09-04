package main

import "fmt"

func main() {

	// creating a pointer

	var b = 10

	var a *int

	fmt.Println(a)

	a = &b

	fmt.Println("Value of a is ", a)
	fmt.Println("Value stored at a is ", *a)

	*a = *a * 2

	fmt.Println("New value of b : ", b)

}
