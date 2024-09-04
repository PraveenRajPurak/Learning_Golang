package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a number : ")

	input, _ := reader.ReadString('\n')

	fmt.Println(input)

	// converting the input to int and adding 5 to it.

	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Increased Value = ", num+5)

}
