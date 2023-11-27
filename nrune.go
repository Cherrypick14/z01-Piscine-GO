package piscine

func NRune(s string, n int) rune {
	var nth rune

	for i, c := range s {
		if i+1 == n {
			nth = c
		}
	}
	return nth
}
