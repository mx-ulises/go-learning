func compressedString(word string) string {
	output := []rune{}
	current, count := rune(word[0]), 0
	for _, c := range word {
		if c == current && count < 9 {
			count++
		} else {
			output = append(output, '0'+rune(count))
			output = append(output, current)
			current = c
			count = 1
		}
	}
	output = append(output, '0'+rune(count))
	output = append(output, current)
	return string(output)
}
