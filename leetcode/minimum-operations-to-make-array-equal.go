func minOperations(n int) int {
	output, i := 0, 1
	for i < n {
		output += n - i
		i += 2
	}
	return output
}
