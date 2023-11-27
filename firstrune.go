package piscine

func FirstRune(s string) rune {
	for _, c := range s {
		return c
	}
	return rune(0)
}
