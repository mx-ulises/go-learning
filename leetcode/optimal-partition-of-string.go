func partitionString(s string) int {
	charsInSubstring := make([]int, 26)
	output := 1
	for i := 0; i < len(s); i++ {
		c := s[i] - 97
		if charsInSubstring[c] == output {
			output += 1
		}
		charsInSubstring[c] = output
	}
	return output
}
