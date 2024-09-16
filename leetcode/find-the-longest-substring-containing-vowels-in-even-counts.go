func findTheLongestSubstring(s string) int {
	vowelMap := map[byte]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
	maskMap := map[int]int{31: -1}
	current, n := 31, len(s)
	maximal := 0
	for i := 0; i < n; i++ {
		current ^= vowelMap[s[i]]
		position, ok := maskMap[current]
		if ok {
			maximal = max(maximal, i-position)
		} else {
			maskMap[current] = i
		}
	}
	return maximal
}
