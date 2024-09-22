func stableMountains(height []int, threshold int) []int {
	output := []int{}
	n := len(height) - 1
	for i := 0; i < n; i++ {
		if threshold < height[i] {
			output = append(output, i+1)
		}
	}
	return output
}
