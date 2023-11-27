package main

import (
	"fmt"
	"os"
)

func Separate(s string, sep string) {
	if len(s) == 0 || len(sep) == 0 {
		return
	}
	newStr := ""
	length := len(sep)
	found := false

	for i := 0; i < len(s); i++ {
		if i+length <= len(s) && s[i:i+length] == sep {
			newStr += " "
			i += length - 1
			found = true
		} else {
			newStr += string(s[i])
		}
	}

	if found {
		fmt.Println(newStr)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		return
	}

	Separate(args[0], args[1])
}
