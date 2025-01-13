func minimumLength(s string) int {
	charCount := map[rune]int{}
	for _, c := range s {
		charCount[c]++
	}
	output := 0
	for _, count := range charCount {
		if count&1 == 1 {
			output += 1
		} else {
			output += 2
		}
	}
	return output
}
