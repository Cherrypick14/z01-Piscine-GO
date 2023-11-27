package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascend := true
	descend := true
	for i := 0; i < len(a)-1; i++ {

		if f(a[i], a[i+1]) > 0 {
			descend = false
		}
		if f(a[i], a[i+1]) < 0 {
			ascend = false
		}
	}
	return ascend || descend
}
