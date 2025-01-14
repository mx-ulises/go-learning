func countBits(n int) []int {
	output := []int{0}
	i := 0
	for i < n {
		j, m := 0, len(output)
		for j < m && i < n {
			output = append(output, output[j]+1)
			i++
			j++
		}
	}
	return output
}
