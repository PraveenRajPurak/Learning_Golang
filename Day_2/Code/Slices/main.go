package main

import (
	"fmt"
	"sort"
)

func main() {

	// Slices are dynamic arrays in go
	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println("arr : ", arr)

	var arr2 = []int{6, 7, 8, 9, 10}
	fmt.Println("arr2 : ", arr2)

	arr3 := []int{11, 12, 13, 14, 15}
	fmt.Println("arr3 : ", arr3)

	// Using make keyword

	var arr4 = make([]int, 5)
	arr4[0] = 1
	arr4[1] = 2
	arr4[2] = 3
	arr4[3] = 4
	arr4[4] = 5
	fmt.Println("arr4 : ", arr4)

	// append

	arr4 = append(arr4, 6, 7, 8, 9, 10)
	fmt.Println("arr4 : ", arr4)

	arr3 = append(arr3[1:4])
	fmt.Println("arr3 : ", arr3)

	arr2 = append(arr2[:4])
	fmt.Println("arr2 : ", arr2)

	// Sorting

	fmt.Println("arr before sorting : ", arr)
	arr = append(arr, 4, 25, 1, 2)
	sort.Ints(arr)
	fmt.Println("arr after sorting : ", arr)

	arr5 := []int{20, 21, 22, 23, 24, 25}
	fmt.Println("arr5 : ", arr5)

	// deleting value of a particular index
	index := 2
	arr5 = append(arr5[:index], arr5[index+1:]...)
	fmt.Println("New arr5 : ", arr5)

}
