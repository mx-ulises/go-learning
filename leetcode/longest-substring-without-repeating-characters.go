func lengthOfLongestSubstring(s string) int {
	max_length := 0
	start := 0
	characters := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		new_start, is_duplicate_char := characters[c]
		if is_duplicate_char && start <= new_start {
			start = new_start + 1
		}
		characters[c] = i
		length := i - start + 1
		if max_length < length {
			max_length = length
		}
	}
	return max_length
}
