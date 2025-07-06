package leetcode

func transpose(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	transposed := make([][]int, m)
	for j := 0; j < m; j++ {
		transposed[j] = make([]int, n)
		for i := 0; i < n; i++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
