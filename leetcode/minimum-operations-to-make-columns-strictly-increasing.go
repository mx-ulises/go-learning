func minimumOperations(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	operations := 0
	for j := 0; j < m; j++ {
		current := grid[0][j] + 1
		for i := 1; i < n; i++ {
			operations += max(0, current-grid[i][j])
			current = max(current, grid[i][j]) + 1
		}
	}
	return operations
}
