func to_upper(c rune) rune {
	if 'a' <= c && c <= 'z' {
		return c - 32
	}
	return c
}

func to_lower(c rune) rune {
	if 'A' <= c && c <= 'Z' {
		return c + 32
	}
	return c
}

func capitalizeTitle(title string) string {
	output := []rune{}
	currentWordSize := 0
	n := 0
	for _, c := range title {
		if c == ' ' {
			if 2 < currentWordSize {
				j := n - currentWordSize
				output[j] = to_upper(output[j])
			}
			currentWordSize = 0
		} else {
			c = to_lower(c)
			currentWordSize++
		}
		output = append(output, c)
		n++
	}
	if 2 < currentWordSize {
		j := n - currentWordSize
		output[j] = to_upper(output[j])
	}
	return string(output)
}
