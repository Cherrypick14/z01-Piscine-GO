package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}

	for i := nb; ; i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			return i
		}
	}
}
