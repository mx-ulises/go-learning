func minChanges(s string) int {
	n := len(s) / 2
	changesCount := 0
	for i := 0; i < n; i++ {
		j := i * 2
		k := j + 1
		if s[j] != s[k] {
			changesCount++
		}
	}
	return changesCount
}
