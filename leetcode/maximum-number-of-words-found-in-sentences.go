func GetWordCount(sentence string) int {
	wordCount := 1
	for _, c := range sentence {
		if c == ' ' {
			wordCount++
		}
	}
	return wordCount
}

func mostWordsFound(sentences []string) int {
	maxWordCount := 0
	for _, sentence := range sentences {
		wordCount := GetWordCount(sentence)
		if maxWordCount < wordCount {
			maxWordCount = wordCount
		}
	}
	return maxWordCount
}
