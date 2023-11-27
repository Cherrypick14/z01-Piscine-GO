package main

import (
	"fmt"
	"os"
)

// Write a program which separates the words of a string, which puts them in a string array and which then prints it to standard output.

func splitprog(str string, sep string) {
	word := ""
	for i := 0; i < len(str); i++ {
		if len(str)-i >= len(sep) && str[i:i+len(sep)] == sep {
			word += " "
			i = i + len(sep) - 1 // we are jumping the separator
		} else {
			word += string(str[i])
		}
	}
	fmt.Println(word)
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <string> <separator>")
		return
	}
	str := os.Args[1]
	sep := os.Args[2]

	splitprog(str, sep)
}
