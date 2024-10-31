func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, c := range s {
		if c != rune(words[i][0]) {
			return false
		}
	}
	return true
}
