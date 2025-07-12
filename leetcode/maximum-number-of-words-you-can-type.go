package leetcode

func canBeTypedWords(text string, brokenLetters string) int {
	missing := map[rune]bool{}
	for _, c := range brokenLetters {
		missing[c] = true
	}
	validWord, validWordCount := true, 0
	for _, c := range text {
		if c == ' ' {
			if validWord {
				validWordCount++
			}
			validWord = true
		}
		validWord = validWord && !missing[c]
	}
	if validWord {
		validWordCount++
	}
	return validWordCount
}
