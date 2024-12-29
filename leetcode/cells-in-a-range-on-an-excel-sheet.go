func cellsInRange(s string) []string {
	output := []string{}
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			output = append(output, string([]byte{i, j}))
		}
	}
	return output
}
