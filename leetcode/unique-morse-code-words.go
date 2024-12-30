func getTranslatedWord(ALPHABET []string, word string) string {
	morseChars := []string{}
	for i := 0; i < len(word); i++ {
		j := int(word[i] - 'a')
		morseChars = append(morseChars, ALPHABET[j])
	}
	return strings.Join(morseChars, "")
}

func uniqueMorseRepresentations(words []string) int {
	ALPHABET := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	uniqueWords := map[string]bool{}
	for _, word := range words {
		translatedWord := getTranslatedWord(ALPHABET, word)
		uniqueWords[translatedWord] = true
	}
	return len(uniqueWords)
}
