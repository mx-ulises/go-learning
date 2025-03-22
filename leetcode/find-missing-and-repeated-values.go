func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	m := n * n
	visitCount := make([]int, m)
	for _, row := range grid {
		for _, num := range row {
			visitCount[num-1]++
		}
	}
	output := []int{-1, -1}
	for i := 0; i < m; i++ {
		if visitCount[i] == 0 {
			output[1] = i + 1
		}
		if visitCount[i] == 2 {
			output[0] = i + 1
		}
	}
	return output
}
