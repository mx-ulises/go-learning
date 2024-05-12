func GetCharacterSet(word *string) *map[rune]bool {
	set := map[rune]bool{}
	for _, c := range *word {
		set[c] = true
	}
	return &set
}

func IsConsistentWord(allowedSet *map[rune]bool, word *string) bool {
	for _, c := range *word {
		if (*allowedSet)[c] == false {
			return false
		}
	}
	return true
}

func countConsistentStrings(allowed string, words []string) int {
	allowedSet, count := GetCharacterSet(&allowed), 0
	for i := 0; i < len(words); i++ {
		if IsConsistentWord(allowedSet, &(words[i])) {
			count++
		}
	}
	return count
}
