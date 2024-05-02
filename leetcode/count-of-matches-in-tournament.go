func numberOfMatches(n int) int {
	matches := 0
	for 1 < n {
		matches += n / 2
		n = (n / 2) + (n & 1)
	}
	return matches
}
