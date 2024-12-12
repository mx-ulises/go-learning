func countKConstraintSubstrings(s string, k int) int {
	l, r, substrings := 0, 0, 0
	bitcount := map[byte]int{}
	for i := 0; i < len(s); i++ {
		r++
		bitcount[s[i]]++
		for k < bitcount['0'] && k < bitcount['1'] {
			bitcount[s[l]]--
			l++
		}
		substrings += r - l
	}
	return substrings
}
