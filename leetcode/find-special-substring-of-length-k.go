package leetcode

func checkDiff(s string, i int, c byte) bool {
	return i < 0 || i == len(s) || c != s[i]
}

func hasSpecialSubstring(s string, k int) bool {
	count, current := 0, s[0]
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == current {
			count++
		} else {
			count, current = 1, c
		}
		if count == k && checkDiff(s, i-k, c) && checkDiff(s, i+1, c) {
			return true
		}
	}
	return false
}
