func findWordsContaining(words []string, x byte) []int {
	filteredWords := []int{}
	xRune := rune(x)
	for i, word := range words {
		for _, c := range word {
			if c == xRune {
				filteredWords = append(filteredWords, i)
				break
			}
		}
	}
	return filteredWords
}
