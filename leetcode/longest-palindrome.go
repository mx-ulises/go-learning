func longestPalindrome(s string) int {
	charCount := map[rune]int{}
	for _, c := range s {
		charCount[c]++
	}
	length, odds := 0, 0
	for _, count := range charCount {
		length += count - (count & 1)
		odds += count & 1
	}
	return length + min(odds, 1)
}
