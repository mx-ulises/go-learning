func reversePrefix(word string, ch byte) string {
	i, n := 0, len(word)
	output := []byte{}
	for i < n && word[i] != ch {
		i++
	}
	if i == n {
		return word
	}
	for j := i; 0 <= j; j-- {
		output = append(output, word[j])
	}
	i++
	for i < n {
		output = append(output, word[i])
		i++
	}
	return string(output)
}
