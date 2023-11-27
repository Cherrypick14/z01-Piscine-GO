package piscine

func ConcatParams(args []string) string {
	new := ""
	newline := "\n"
	for i, arg := range args {
		new += arg
		if i < len(args)-1 {
			new += newline
		}

	}
	return new
}
