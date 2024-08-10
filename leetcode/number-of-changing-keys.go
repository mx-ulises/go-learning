func countKeyChanges(s string) int {
	s = strings.ToLower(s)
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			count++
		}
	}
	return count
}
