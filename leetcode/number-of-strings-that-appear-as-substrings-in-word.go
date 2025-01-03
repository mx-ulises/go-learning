func isSubString(word, pattern string) bool {
	n, m := len(word), len(pattern)
	for i := 0; i < n; i++ {
		j, k := i, 0
		match := true
		for j < n && k < m {
			if word[j] != pattern[k] {
				match = false
				break
			}
			j++
			k++
		}
		if match && k == m {
			return true
		}
	}
	return false
}

func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, pattern := range patterns {
		if isSubString(word, pattern) {
			count++
		}
	}
	return count
}
