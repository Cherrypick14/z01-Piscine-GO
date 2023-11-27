package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myargs := os.Args[1:]

	// Iterate over the `myargs` slice in reverse order and print each argument in reverse order.
	for i := len(myargs) - 1; i >= 0; i-- {
		for _, j := range myargs[i] {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}
}
