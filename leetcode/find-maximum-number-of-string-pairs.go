func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func maximumNumberOfStringPairs(words []string) int {
	visited := map[string]bool{}
	count := 0
	for _, word := range words {
		if visited[word] {
			count++
		} else {
			visited[Reverse(word)] = true
		}
	}
	return count
}
