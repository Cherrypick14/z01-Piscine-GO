package piscine

func DescendAppendRange(max, min int) []int {
	runes := []int{}

	if max <= min {
		return []int{}
	}
	for i := max; i > min; i-- {
		runes = append(runes, i)
	}
	return runes
}
