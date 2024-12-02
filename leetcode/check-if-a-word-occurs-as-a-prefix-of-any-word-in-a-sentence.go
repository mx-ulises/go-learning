func isPrefixOfWord(sentence string, searchWord string) int {
	currentWord := 1
	searchIndex := 0
	isWordOk := true
	for i, _ := range sentence {
		if sentence[i] == ' ' {
			currentWord++
			searchIndex = 0
			isWordOk = true
		} else if isWordOk && sentence[i] == searchWord[searchIndex] {
			searchIndex++
			if searchIndex == len(searchWord) {
				return currentWord
			}
		} else {
			isWordOk = false
		}
	}
	return -1
}
