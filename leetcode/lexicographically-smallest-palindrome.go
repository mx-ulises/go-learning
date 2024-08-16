func makeSmallestPalindrome(s string) string {
	n := len(s)
	output := make([]byte, n)
	i, j := 0, n-1
	for i <= j {
		c := min(s[i], s[j])
		output[i] = c
		output[j] = c
		i++
		j--
	}
	return string(output)
}
