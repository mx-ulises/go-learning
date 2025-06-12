func percentageLetter(s string, letter byte) int {
	charCount := [26]int{}
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		charCount[c] += 100
	}
	return charCount[int(letter-'a')] / len(s)
}
