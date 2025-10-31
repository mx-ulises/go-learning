package leetcode

func getCharNumber(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a'
	}
	return c - 'A'
}

func toLower(c byte) byte {
	return getCharNumber(c) + 'a'
}

func toUpper(c byte) byte {
	return getCharNumber(c) + 'A'
}

func generateTag(caption string) string {
	start := 0
	for start < len(caption) && caption[start] == ' ' {
		start++
	}
	if start == len(caption) {
		return "#"
	}
	output, upper := []byte{'#', toLower(caption[start])}, false
	for i := start + 1; i < len(caption); i++ {
		if caption[i] != ' ' {
			if upper {
				output, upper = append(output, toUpper(caption[i])), false
			} else {
				output = append(output, toLower(caption[i]))
			}
			if len(output) == 100 {
				break
			}
		} else {
			upper = true
		}
	}
	return string(output)
}
