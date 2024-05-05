func GetWords(s string) (map[rune][]rune, rune) {
	words := map[rune][]rune{}
	current := []rune{}
	maximal := '1'
	for _, c := range s {
		if '1' <= c && c <= '9' {
			words[c] = current
			current = []rune{}
			maximal = max(maximal, c)
		} else if c != ' ' {
			current = append(current, c)
		}
	}
	return words, maximal
}

func BuildOutput(words map[rune][]rune, maximal rune) string {
	output := []rune{}
	output = append(output, words['1']...)
	for c := '2'; c <= maximal; c++ {
		output = append(output, ' ')
		output = append(output, words[c]...)
	}
	return string(output)
}

func sortSentence(s string) string {
	words, maximal := GetWords(s)
	return BuildOutput(words, maximal)
}
