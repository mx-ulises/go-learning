func GetCharCount(word *string) *map[string]int {
	charCount := map[string]int{}
	for _, c := range *word {
		charCount[string(c)]++
	}
	return &charCount
}

func GetCommonCharacters(commonCharCount, charCount *map[string]int) *map[string]int {
	output := map[string]int{}
	for c, _ := range *commonCharCount {
		output[c] = min((*commonCharCount)[c], (*charCount)[c])
	}
	return &output
}

func commonChars(words []string) []string {
	commonCharCount := GetCharCount(&(words[0]))
	for i := 1; i < len(words); i++ {
		charCount := GetCharCount(&(words[i]))
		commonCharCount = GetCommonCharacters(commonCharCount, charCount)
	}
	output := []string{}
	for c, count := range *commonCharCount {
		for i := 0; i < count; i++ {
			output = append(output, c)
		}
	}
	return output
}
