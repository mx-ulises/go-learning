func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	rows, cols := make([]bool, n), make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}
