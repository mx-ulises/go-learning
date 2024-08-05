func kthDistinct(arr []string, k int) string {
	wordCount := map[string]int{}
	for _, word := range arr {
		wordCount[word]++
	}
	for _, word := range arr {
		if wordCount[word] == 1 && k == 1 {
			return word
		}
		if wordCount[word] == 1 {
			k--
		}
	}
	return ""
}
