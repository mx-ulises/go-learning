func minSteps(s string, t string) int {
	charCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charCount[s[i]]++
		charCount[t[i]]--
	}
	steps := 0
	for _, count := range charCount {
		if 0 < count {
			steps += count
		}
	}
	return steps
}
