package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

const url = "https://www.gutenberg.org/files/1342/1342-0.txt"

func main() {
	resp, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response type %T", resp)
    defer resp.Body.Close()
    
    bytes, err := ioutil.ReadAll(resp.Body)
    checkError(err)
    
    content := string(bytes)
    fmt.Print(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
