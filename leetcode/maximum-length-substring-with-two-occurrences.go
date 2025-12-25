package leetcode

func maximumLengthSubstring(s string) int {
	start, maxLen := 0, 0
	charCount := map[byte]int{}
	for end, _ := range s {
		start = checkAndUpdateStart(start, charCount, s, s[end])
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func checkAndUpdateStart(start int, charCount map[byte]int, s string, endChar byte) int {
	charCount[endChar]++
	for 2 < charCount[endChar] {
		startChar := s[start]
		charCount[startChar]--
		start++
	}
	return start
}
