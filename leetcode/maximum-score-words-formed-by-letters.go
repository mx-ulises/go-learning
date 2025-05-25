func getWordScore(word string, score []int) int {
	wordScore := 0
	for i := 0; i < len(word); i++ {
		wordScore += score[int(word[i]-'a')]
	}
	return wordScore
}

func getWordScores(words []string, score []int) map[string]int {
	wordScores := map[string]int{}
	for _, word := range words {
		wordScores[word] = getWordScore(word, score)
	}
	return wordScores
}

func modifyCountFromWord(letterCount map[byte]int, word string, add int) bool {
	valid := true
	for i := 0; i < len(word); i++ {
		letterCount[word[i]] += add
		if letterCount[word[i]] < 0 {
			valid = false
		}
	}
	return valid
}

func getMaxScore(words []string, letterCount map[byte]int, wordScores map[string]int, i, current int) int {
	if i == len(words) {
		return current
	}
	word := words[i]
	i++
	output := current
	if modifyCountFromWord(letterCount, word, -1) {
		output = max(output, getMaxScore(words, letterCount, wordScores, i, current+wordScores[word]))
	}
	modifyCountFromWord(letterCount, word, 1)
	return max(output, getMaxScore(words, letterCount, wordScores, i, current))
}

func getLetterCount(letters []byte) map[byte]int {
	letterCount := map[byte]int{}
	for _, letter := range letters {
		letterCount[letter]++
	}
	return letterCount
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	wordScores := getWordScores(words, score)
	letterCount := getLetterCount(letters)
	return getMaxScore(words, letterCount, wordScores, 0, 0)
}
