package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		if len(args) > 1 {
			fmt.Println("Too many arguments")
		} else {
			fmt.Println("File name missing")
		}
		return
	}
	filename := args[0]
	ReadFile(filename)
}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
