func isPrefixSuffixPair(str1, str2 string) bool {
	n1, n2 := len(str1), len(str2)
	return n1 <= n2 && str1 == str2[:n1] && str1 == str2[n2-n1:]
}

func countPrefixSuffixPairs(words []string) int {
	n := len(words)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixSuffixPair(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}
