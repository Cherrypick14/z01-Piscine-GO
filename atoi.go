package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
			continue
		} else if i == 0 && char == '+' {
			continue
		} else if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
