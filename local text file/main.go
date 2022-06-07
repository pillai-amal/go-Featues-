package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "I am writing this sentence into text file"
	file, err := os.Create("./textfile.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote the the filw with %v characters", length)
	defer file.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
