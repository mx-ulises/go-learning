func lengthOfLastWord(s string) int {
	currentWordLen := 0
	newWord := true
	for _, c := range s {
		if c == ' ' {
			newWord = true
		} else {
			if newWord == true {
				currentWordLen = 0
				newWord = false
			}
			currentWordLen++
		}
	}
	return currentWordLen
}
