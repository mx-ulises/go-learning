func GetValue(matrix *[][]int, i int, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return (*matrix)[i][j]
}

func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				prev_i := GetValue(&matrix, i-1, j)
				prev_j := GetValue(&matrix, i, j-1)
				prev := GetValue(&matrix, i-1, j-1)
				matrix[i][j] += min(prev_i, prev_j, prev)
				count += matrix[i][j]
			}
		}
	}
	return count
}
