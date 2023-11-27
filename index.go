package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	c := toFind[0]
	for i := 0; i < len(s); i++ {
		if s[i] == c && len(s)-i >= len(toFind) && s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
