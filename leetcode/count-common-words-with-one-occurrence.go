package leetcode

func countWords(words1 []string, words2 []string) int {
	wordCount := map[string]int{}
	for _, word := range words1 {
		wordCount[word]++
	}
	for _, word := range words2 {
		wordCount[word] += 1000
	}
	total := 0
	for _, count := range wordCount {
		if count == 1001 {
			total++
		}
	}
	return total
}
