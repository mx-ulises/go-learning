func IsPalindrome(word string) bool {
	n := len(word)
	m := n / 2
	for i := 0; i < m; i++ {
		if word[i] != word[n-i-1] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) string {
	for _, word := range words {
		if IsPalindrome(word) {
			return word
		}
	}
	return ""
}
