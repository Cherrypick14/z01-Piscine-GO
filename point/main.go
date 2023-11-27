package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func PrintStr(s string) {
	arraystr := []rune(s)
	for i := range arraystr {
		z01.PrintRune(arraystr[i])
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintValue(a int) {
	r := '0'
	if a/10 > 0 {
		PrintValue(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func main() {
	var points point

	setPoint(&points)
	PrintStr("x = ")
	PrintValue(points.x)
	PrintStr(", y = ")
	PrintValue(points.y)
	z01.PrintRune('\n')
}
