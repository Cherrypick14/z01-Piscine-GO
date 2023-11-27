package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(args); j++ {
			if args[j] < args[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			args[i], args[minIndex] = args[minIndex], args[i]
		}
	}

	for _, arg := range args {
		// Convert the string to a slice of runes.
		runes := []rune(arg)

		// Iterate over the runes and print them one at a time.
		for _, rune := range runes {
			z01.PrintRune(rune)
		}

		// Print a newline character.
		z01.PrintRune('\n')
	}
}
