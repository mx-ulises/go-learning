func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		c := s[left]
		for left <= right && c == s[left] {
			left++
		}
		for left <= right && c == s[right] {
			right--
		}
	}
	return right - left + 1
}
