package piscine

func Compact(ptr *[]string) int {
	count := 0
	store := *ptr

	for _, char := range store {
		if char != "" {
			store[count] = char
			count++
		}
	}
	*ptr = store[:count]
	return count
}
