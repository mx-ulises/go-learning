func IsNewLand(grid *[][]byte, neighbors *[][]int, i, j, n, m int) bool {
	if i < 0 || j < 0 || i == n || j == m || (*grid)[i][j] != '1' {
		return false
	}
	(*grid)[i][j] = '2'
	for _, neighbor := range *neighbors {
		IsNewLand(grid, neighbors, i+neighbor[0], j+neighbor[1], n, m)
	}
	return true
}

func numIslands(grid [][]byte) int {
	count := 0
	neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if IsNewLand(&grid, &neighbors, i, j, n, m) {
				count++
			}
		}
	}
	return count
}
