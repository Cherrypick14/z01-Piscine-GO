package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myargs := os.Args[1:]
	for _, char := range myargs {
		for i := 0; i < len(char); i++ {
			z01.PrintRune(rune(char[i]))
		}
		z01.PrintRune('\n')
	}
}
