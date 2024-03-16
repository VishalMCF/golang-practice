package main

import (
	"fmt"
	"log"
	"os"
)

var file_name = "sample_text.txt"

func createFile() {
	_, err := os.Create(file_name)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello Basic IO..")
	file, err := os.Stat(file_name)
	if err!=nil
	fmt.Println("File was successfully created")
	contents, err := os.ReadFile("sample_text.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
	defer file.Close()
}
