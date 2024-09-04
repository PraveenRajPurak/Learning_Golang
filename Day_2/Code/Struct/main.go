package main

import "fmt"

func main() {

	user1 := User{"Praveen", 22}
	fmt.Println(user1)

	//print a more detailed struct :

	fmt.Printf("Struct value : %+v \n", user1)

}

type User struct {
	Name string
	Age  int
}
