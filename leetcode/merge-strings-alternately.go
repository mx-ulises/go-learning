func mergeAlternately(word1 string, word2 string) string {
	n := min(len(word1), len(word2))
	output := make([]byte, n*2)
	for i := 0; i < n; i++ {
		j := i << 1
		output[j] = word1[i]
		output[j+1] = word2[i]
	}
	output = append(output, word1[n:]...)
	output = append(output, word2[n:]...)
	return string(output)
}
