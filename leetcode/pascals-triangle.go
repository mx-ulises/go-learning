func generate(numRows int) [][]int {
	output := make([][]int, numRows)
	output[0] = []int{1}
	for i := 1; i < numRows; i++ {
		output[i] = make([]int, i+1)
		output[i][0] = 1
		prev := 1
		for j := 1; j < i; j++ {
			output[i][j] = prev + output[i-1][j]
			prev = output[i-1][j]
		}
		output[i][i] = 1
	}
	return output
}
