package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// File handling

	//Creating a file

	file, err := os.Create("./file.txt")

	FileError(err)

	defer file.Close()

	//Writing to a file

	_, err = FileWriter("Hello World", "./file.txt")
	FileError(err)

	//Reading from a file

	data, err := FileReader("./file.txt")
	fmt.Println(data)
	FileError(err)
}
func FileWriter(Content string, file_name string) (int, error) {
	opened_file, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	length, err := io.WriteString(opened_file, Content)

	return length, err
}

func FileReader(File string) (string, error) {
	dataByte, err := os.ReadFile(File)

	return string(dataByte), err
}

func FileError(err error) {
	if err != nil {
		panic(err)
	}
}
