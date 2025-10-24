package leetcode

func getCharCount(s string) map[rune]int {
	charCount := map[rune]int{}
	for _, c := range s {
		charCount[c]++
	}
	return charCount
}

func rearrangeCharacters(s string, target string) int {
	sCharCount := getCharCount(s)
	targetCharCount := getCharCount(target)
	copies := 1000
	for c, count := range targetCharCount {
		copies = min(copies, sCharCount[c]/count)
	}
	return copies
}
