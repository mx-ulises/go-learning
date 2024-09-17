func AddWords(wordCount *map[string]int, s *string) {
	start := 0
	for i, c := range *s {
		if c == ' ' {
			(*wordCount)[(*s)[start:i]]++
			start = i + 1
		}
	}
	(*wordCount)[(*s)[start:]]++
}

func uncommonFromSentences(s1 string, s2 string) []string {
	wordCount := map[string]int{}
	AddWords(&wordCount, &s1)
	AddWords(&wordCount, &s2)
	output := []string{}
	for word, count := range wordCount {
		if count == 1 && 0 < len(word) {
			output = append(output, word)
		}
	}
	return output
}
