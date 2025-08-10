package leetcode

func numberOfSpecialChars(word string) int {
	specialChars := map[rune]int{}
	for _, c := range word {
		if 'A' <= c && c <= 'Z' {
			specialChars[c-'A'] |= 1
		}
		if 'a' <= c && c <= 'z' {
			specialChars[c-'a'] |= 2
		}
	}
	count := 0
	for _, mask := range specialChars {
		if mask == 3 {
			count++
		}
	}
	return count
}
