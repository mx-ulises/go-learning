func countGoodSubstrings(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}
	a, b, count := s[0], s[1], 0
	for i := 2; i < n; i++ {
		c := s[i]
		if a != b && a != c && b != c {
			count++
		}
		a, b = b, c
	}
	return count
}
