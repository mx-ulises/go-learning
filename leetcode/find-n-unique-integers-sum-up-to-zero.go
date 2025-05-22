func sumZero(n int) []int {
	output := make([]int, n)
	i, current := n&1, 1
	for i < n {
		output[i], output[i+1] = current, -current
		i += 2
		current++
	}
	return output
}
