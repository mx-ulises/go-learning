func findWordsContaining(words []string, x byte) []int {
	output := []int{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] == x {
				output = append(output, i)
				break
			}
		}
	}
	return output
}
