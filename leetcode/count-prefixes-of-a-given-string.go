func isPrefix(word, prefix string) bool {
	if len(word) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if word[i] != prefix[i] {
			return false
		}
	}
	return true
}

func countPrefixes(words []string, s string) int {
	prefixes := 0
	for _, word := range words {
		if isPrefix(s, word) {
			prefixes++
		}
	}
	return prefixes
}
