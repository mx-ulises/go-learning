package leetcode

func isPrefixString(s string, words []string) bool {
	i := 0
	for _, word := range words {
		for j := 0; j < len(word); j++ {
			if i == len(s) || s[i] != word[j] {
				return false
			}
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return i == len(s)
}
