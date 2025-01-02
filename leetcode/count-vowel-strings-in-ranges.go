func isVowelString(word string, vowels map[byte]bool) int {
	if vowels[word[0]] && vowels[word[len(word)-1]] {
		return 1
	}
	return 0
}

func getVowelStringCount(words []string, vowels map[byte]bool) []int {
	n := len(words)
	vowelStringCount := make([]int, n+1)
	for i := 0; i < n; i++ {
		vowelStringCount[i+1] = vowelStringCount[i] + isVowelString(words[i], vowels)
	}
	return vowelStringCount
}

func vowelStrings(words []string, queries [][]int) []int {
	VOWELS := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	vowelStringCount := getVowelStringCount(words, VOWELS)
	output := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		output[i] = vowelStringCount[r+1] - vowelStringCount[l]
	}
	return output
}
