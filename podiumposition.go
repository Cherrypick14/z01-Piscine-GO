package piscine

func PodiumPosition(podium [][]string) [][]string {
	result := podium

	for j := 0; j < len(result); j++ {
		for i := 0; i < len(result)-1; i++ {
			if result[i][0] > result[i+1][0] {
				temp1 := result[i]
				result[i] = result[i+1]
				result[i+1] = temp1

			}
		}
	}
	return result
}
