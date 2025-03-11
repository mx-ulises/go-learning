func numberOfSubstrings(s string) int {
	charCount, chars := [3]int{0, 0, 0}, 0
	l, r, n := 0, 0, len(s)
	count := 0
	for r < n {
		for r < n && chars < 3 {
			c := int(s[r] - 'a')
			charCount[c]++
			if charCount[c] == 1 {
				chars++
			}
			r++
		}
		extra := n - r + 1
		for l < r && chars == 3 {
			count += extra
			c := int(s[l] - 'a')
			charCount[c]--
			if charCount[c] == 0 {
				chars--
			}
			l++
		}
	}
	return count
}
