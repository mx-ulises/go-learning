func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	output := make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, m)
		for j := 0; j < m; j++ {
			value := Min(rowSum[i], colSum[j])
			output[i][j] = value
			rowSum[i] -= value
			colSum[j] -= value
		}
	}
	return output
}
