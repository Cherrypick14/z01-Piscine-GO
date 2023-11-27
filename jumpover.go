package piscine

func JumpOver(str string) string {
	if len(str) < 3 || str == "" {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
