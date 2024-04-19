func scoreOfString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		a, b := int(s[i-1]), int(s[i])
		score += max(a, b) - min(a, b)
	}
	return score
}
