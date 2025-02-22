func checkIfPangram(sentence string) bool {
	charCount := 0
	charInWord := make([]bool, 26)
	for i := 0; i < len(sentence); i++ {
		c := int(sentence[i] - 'a')
		if charInWord[c] == false {
			charCount++
			charInWord[c] = true
		}
	}
	return charCount == 26
}
