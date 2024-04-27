func truncateSentence(s string, k int) string {
	output := []byte{}
	i, spaceCount := 0, 0
	for i < len(s) {
		c := s[i]
		if c == ' ' {
			spaceCount++
			if spaceCount == k {
				break
			}
		}
		output = append(output, c)
		i++
	}
	return string(output)
}
