package main

import "fmt"

func main() {

	// Creating maps

	var languages = make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"

	fmt.Println(languages)

	//accessing map values

	fmt.Println("Fullform of JS : ", languages["JS"])

	// Deleting a map value

	delete(languages, "RB")
	fmt.Println(languages)

}
