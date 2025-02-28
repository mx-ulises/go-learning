func areOccurrencesEqual(s string) bool {
	charCount := map[rune]int{}
	maximal := 0
	for _, c := range s {
		charCount[c]++
		maximal = max(maximal, charCount[c])
	}
	for _, count := range charCount {
		if count != maximal {
			return false
		}
	}
	return true
}
