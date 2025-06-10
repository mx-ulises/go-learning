func maxDifference(s string) int {
	charCount := map[byte]int{}
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
	}
	a1, a2 := 0, len(s)
	for _, count := range charCount {
		if count&1 == 0 {
			a2 = min(a2, count)
		} else {
			a1 = max(a1, count)
		}
	}
	return a1 - a2
}
