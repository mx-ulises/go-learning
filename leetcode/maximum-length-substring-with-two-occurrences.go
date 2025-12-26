package leetcode

func maximumLengthSubstring(s string) int {
	start, maxLen := 0, 0
	charCount := map[byte]int{}
	for end, _ := range s {
		start, charCount = checkAndUpdateStartAndCharCount(start, charCount, s, s[end])
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func checkAndUpdateStartAndCharCount(start int, charCount map[byte]int, s string, endChar byte) (int, map[byte]int) {
	charCount[endChar]++
	for 2 < charCount[endChar] {
		start, charCount = updateStartAndCharCount(start, charCount, s[start])
	}
	return start, charCount
}

func updateStartAndCharCount(start int, charCount map[byte]int, startChar byte) (int, map[byte]int) {
	charCount[startChar]--
	return start + 1, charCount
}
