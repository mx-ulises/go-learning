func splitWordsBySeparator(words []string, separator byte) []string {
	output := []string{}
	for _, word := range words {
		start, n := 0, len(word)
		for end := 0; end < n; end++ {
			if word[end] == separator {
				if start < end {
					output = append(output, word[start:end])
				}
				start = end + 1
			}
		}
		if start < n {
			output = append(output, word[start:n])
		}
	}
	return output
}
