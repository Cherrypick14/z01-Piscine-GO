package main

import (
	"fmt"
	"os"
)

func searchnreplace(str, f, re string) string {
	newstr := ""
	for _, char := range str {
		for _, char2 := range f {
			if char == char2 {
				newstr += re
			} else {
				newstr += string(char)
			}
		}
	}
	return newstr
}

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		fmt.Println("Only 3 args allowed")
		return
	}

	fmt.Println(searchnreplace(args[0], args[1], args[2]))
}
