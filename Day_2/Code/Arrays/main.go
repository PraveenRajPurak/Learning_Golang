package main

import "fmt"

func main() {

	// Arrays have fixed length in go

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr : ", arr)

	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println("arr2 : ", arr2)

	var arr3 = [5]int{11, 12, 13, 14, 15}
	fmt.Println("arr3 : ", arr3)
}
