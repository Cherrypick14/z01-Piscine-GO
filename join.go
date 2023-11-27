package piscine

func Join(strs []string, sep string) string {
	result := strs[0]
	for _, char := range strs[1:] {
		result = result + sep + char
	}
	return result
}
