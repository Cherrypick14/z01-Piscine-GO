package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, c := range s {
		digit := int(c - '0')
		result = result*10 + digit
	}
	return result
}
