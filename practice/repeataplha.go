package main

import (
	"fmt"
	"os"
)

func FindIndex(char byte) int {
	count := 0
	if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
		if char >= 'a' && char <= 'z' {
			for i := 'a'; i <= 'z'; i++ {
				count++
				if i == rune(char) {
					return count
				}
			}
		}
		if char >= 'A' && char <= 'Z' {
			for i := 'A'; i <= 'Z'; i++ {
				count++
				if i == rune(char) {
					return count
				}
			}
		}
	}
	return 0
}

func PrintIndex(str string) {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z') {
			fmt.Print(string(str[i]))
		}
		index := FindIndex(str[i])
		for j := 0; j < index; j++ {
			fmt.Print(string(str[i]))
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	PrintIndex(args[0])
	fmt.Println()
}
