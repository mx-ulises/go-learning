func countServers(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rows, cols := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	connected := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && (2 <= rows[i] || 2 <= cols[j]) {
				connected++
			}
		}
	}
	return connected
}
