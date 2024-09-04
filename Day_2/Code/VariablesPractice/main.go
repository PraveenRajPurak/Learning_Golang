package main

import "fmt"

func main() {

	var name string = "Praveen Raj"
	fmt.Println(name)
	fmt.Printf("Type : %T \n", name)

	var age int = 22
	fmt.Println(age)
	fmt.Printf("Type : %T \n", age)

	var randomNo uint8 = 255
	fmt.Println(randomNo)
	fmt.Printf("Type : %T \n", randomNo)

	var randomNo2 uint16 = 65535
	fmt.Println(randomNo2)
	fmt.Printf("Type : %T \n", randomNo2)

	var grade float32 = 8.5
	fmt.Println(grade)
	fmt.Printf("Type : %T \n", grade)

	var largeFloat float64 = 8.532323232323
	fmt.Println(largeFloat)
	fmt.Printf("Type : %T \n", largeFloat)

	var isStudent bool = true
	fmt.Println(isStudent)
	fmt.Printf("Type : %T \n", isStudent)

	// aliases

	var smallNo byte = 34
	fmt.Println(smallNo)
	fmt.Printf("Type : %T \n", smallNo)

	var smallNo2 rune = 42213
	fmt.Println(smallNo2)
	fmt.Printf("Type : %T \n", smallNo2)

	// implicit type

	var Surname = "Raj"
	fmt.Println(Surname)
	fmt.Printf("Type : %T \n", Surname)

	// no var

	institute := "JIIT"
	fmt.Println(institute)
	fmt.Printf("Type : %T \n", institute)

	// constants

	const value int = 5
	fmt.Println(value)
	fmt.Printf("Type : %T \n", value)
}
