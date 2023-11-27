package piscine

func AlphaCount(s string) int {
	counter := 0

	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			counter++
		}
	}
	return counter
}
