package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(num int) {
	if num == 0 {
		z01.PrintRune('0')
	}
	max := 0

	temp := num

	for temp != 0 {
		digit := temp % 10
		if digit > max {
			max = digit
		}
		temp /= 10
	}

	list := make([]int, max+1)

	for num != 0 {
		digit := num % 10
		list[digit]++
		num /= 10
	}

	for i := 0; i <= max; i++ {
		for j := 0; j < list[i]; j++ {
			z01.PrintRune(rune(i + '0'))
		}
	}
}
