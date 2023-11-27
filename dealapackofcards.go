package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Manlen(slice []int) int {
	length := 0
	for range slice {
		length++
	}
	return length
}

func DealAPackOfCards(deck []int) {
	length := Manlen(deck)
	val := 0
	for i := 1; i <= length/3; i++ {
		fmt.Printf("Player ")
		fmt.Printf("%v", i)
		fmt.Printf(": ")
		for j := val; j < val+3; j++ {
			fmt.Printf("%v", deck[j])
			if j < val+2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		val += 3
	}
}
