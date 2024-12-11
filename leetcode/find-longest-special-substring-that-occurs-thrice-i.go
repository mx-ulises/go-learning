func getSpecialSubtringMap(s string) map[byte][]int {
	specialSubstrings := map[byte][]int{}
	specialSubstrings[s[0]] = append(specialSubstrings[s[0]], 1)
	currentCount := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			currentCount = 0
		}
		if len(specialSubstrings[s[i]]) == currentCount {
			specialSubstrings[s[i]] = append(specialSubstrings[s[i]], 0)
		}
		specialSubstrings[s[i]][currentCount]++
		currentCount++
	}
	return specialSubstrings
}

func maximumLength(s string) int {
	specialSubstrings = getSpecialSubtringMap(s)
	maximal := -1
	for _, lengths := range specialSubstrings {
		n := len(lengths) - 1
		stringCount := 0
		for i := n; max(0, maximal) <= i; i-- {
			stringCount += lengths[i]
			if 3 <= stringCount {
				maximal = i + 1
			}
		}
	}
	return maximal
}
