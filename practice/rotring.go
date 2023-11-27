package main

import (
	"fmt"
	"os"
)

func RotateStr(str string) string {
	if len(str) == 0 {
		return "" // Return an empty string if input is empty
	}
	first := ""
	start := true
	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && start {
			start = false
		}
		if !start {
			if string(str[i]) != " " {
				first += string(str[i])
			} else {
				break
			}
		}
	}
	rest := ""
	prev := ""
	for i := range str[len(first):] {
		if string(str[len(first)+i]) == " " && prev == " " {
			continue
		}
		rest += string(str[len(first)+i])
		prev = string(str[len(first)+i])
	}
	rest += " "
	return rest + first
}

func main() {
	args := os.Args[1:]
	stri := ""
	for _, char := range args {
		for i := 0; i < len(char); i++ {
			stri += string(char[i])
		}
	}
	fmt.Println(RotateStr(stri))
}
