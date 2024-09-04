package main

import (
	"fmt"
	"math/rand"
)

func main() {

	value := 20

	// if else

	if value < 10 {
		fmt.Println("value is less than 10")
	} else if value > 10 {
		fmt.Println("value is greater than 10")
	} else {
		fmt.Println("value is equal to 10")
	}

	// switch case

	random := rand.Intn(6) + 1
	fmt.Println("random : ", random)

	switch random {
	case 1:
		fmt.Println("random is 1")

	case 2:
		fmt.Println("random is 2")

	case 3:
		fmt.Println("random is 3")

	case 4:
		fmt.Println("random is 4")

	case 5:
		fmt.Println("random is 5")

	case 6:
		fmt.Println("random is 6")

	default:
		fmt.Println("???")
	}

	// For Loops

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("One way : ")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	fmt.Println("Second way : ")
	for index := range values {
		fmt.Println(index)
	}

	fmt.Println("Third way : ")
	for index, value := range values {
		fmt.Printf("Value at %v is %v \n", index, value)
	}

	val := 1
	fmt.Println("Fourth way : ")
	for val < 5 {
		fmt.Println(val)
		val++
	}

}
