package main

import "piscine"

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
