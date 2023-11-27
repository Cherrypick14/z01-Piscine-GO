package main

import (
	"fmt"
	"os"
)

func runeToWord(r []rune) string {
	str := ""
	for _, s := range r {
		str += string(s)
	}
	return str
}

func RotString(s string) {
	str := s + " "
	words := []string{}
	runes := []rune{}

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i+1] != ' ' {
			runes = append(runes, rune(str[i]))
		} else {
			word := runeToWord(runes)
			words = append(words, word)
			runes = nil
		}
	}
	fmt.Println(words)
}

func main() {
	args := os.Args[1:]
	RotString(args[0])
}
