package piscine

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	return num[len(num)/2]
}
