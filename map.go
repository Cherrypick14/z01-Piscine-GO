package piscine

func Map(f func(int) bool, a []int) []bool {
	results := make([]bool, len(a))
	for i, v := range a {
		results[i] = f(v)
	}
	return results
}
