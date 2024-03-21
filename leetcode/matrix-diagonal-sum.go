func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	j := n - 1
	for i := 0; i < n; i++ {
		sum += mat[i][i] + mat[i][j]
		if i == j {
			sum -= mat[i][i]
		}
		j--
	}
	return sum
}
