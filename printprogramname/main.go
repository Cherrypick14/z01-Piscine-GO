package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	newargs := os.Args[0]
	for i, char := range newargs {
		if i > 1 {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
