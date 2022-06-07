// This is a simple go program to demostrate how to
// write a simple text file to working directory
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "I am writing this sentence into text file"
	file, err := os.Create("./textfile.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote the the filw with %v characters \n", length)
	defer file.Close()
	defer readFile("textfile.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Printf("The %v contains : %v", fileName, string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
