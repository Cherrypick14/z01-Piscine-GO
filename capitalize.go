package piscine

func IsAlphaNumeric(char rune) bool {
	return (char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9')
}

func Capitalize(s string) string {
	output := []rune(s)
	toCapitalize := true

	for i, char := range output {
		if IsAlphaNumeric(char) {
			if toCapitalize {
				if char >= 'a' && char <= 'z' {
					output[i] = char - 'a' + 'A'
				}
				toCapitalize = false
			} else {
				if char >= 'A' && char <= 'Z' {
					output[i] = char - 'A' + 'a'
				}
			}
		} else {
			toCapitalize = true
		}
	}
	return string(output)
}
