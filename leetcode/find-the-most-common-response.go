func findCommonResponse(responses [][]string) string {
	wordExistencePerDay := map[string]int{}
	for _, daySurvery := range responses {
		viewed := map[string]bool{}
		for _, word := range daySurvery {
			if !viewed[word] {
				wordExistencePerDay[word]++
				viewed[word] = true
			}
		}
	}
	mostCommonWord, mostCommonWordCount := "", 0
	for word, count := range wordExistencePerDay {
		if mostCommonWordCount < count {
			mostCommonWordCount = count
			mostCommonWord = word
		} else if mostCommonWordCount == count {
			mostCommonWord = min(mostCommonWord, word)
		}
	}
	return mostCommonWord
}
