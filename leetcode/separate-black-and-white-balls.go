func minimumSteps(s string) int64 {
	l, r := 0, len(s)-1
	swaps := 0
	for l < r {
		for l < r && s[l] == '0' {
			l++
		}
		for l < r && s[r] == '1' {
			r--
		}
		swaps += r - l
		l++
		r--
	}
	return int64(swaps)
}
