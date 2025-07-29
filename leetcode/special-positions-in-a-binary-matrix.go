package leetcode

func numSpecial(mat [][]int) int {
	n, m := len(mat), len(mat[0])
	rows, cols := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rows[i] == 1 && cols[j] == 1 && mat[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
