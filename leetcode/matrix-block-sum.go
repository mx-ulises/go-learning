func getBlockSum(mat [][]int, x, y, k int) int {
	blockSum := 0
	startRow, endRow := max(0, x-k), min(len(mat)-1, x+k)
	startCol, endCol := max(0, y-k), min(len(mat[0])-1, y+k)
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			blockSum += mat[i][j]
		}
	}
	return blockSum
}

func initMatrix(n, m int) [][]int {
	output := make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, m)
	}
	return output
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	n, m := len(mat), len(mat[0])
	output := initMatrix(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			output[i][j] = getBlockSum(mat, i, j, k)
		}
	}
	return output
}
