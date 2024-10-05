func GetChars(s *string) *map[rune]int {
	charCount := map[rune]int{}
	for _, c := range *s {
		charCount[c]++
	}
	return &charCount
}

func checkInclusion(s1 string, s2 string) bool {
	chars, charCount := GetChars(&s1), len(s1)
	charsInBuffer := map[rune]int{}
	bufferCharCount := 0
	buffer := []rune{}
	for _, c := range s2 {
		buffer = append(buffer, c)
		charsInBuffer[c]++
		bufferCharCount++
		for (*chars)[c] < charsInBuffer[c] {
			d := buffer[0]
			buffer = buffer[1:]
			charsInBuffer[d]--
			bufferCharCount--
		}
		if bufferCharCount == charCount {
			return true
		}
	}
	return false
}
