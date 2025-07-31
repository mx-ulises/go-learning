package leetcode

func projectionArea(grid [][]int) int {
	n := len(grid)
	xy, yz, zx := 0, 0, 0
	for i := 0; i < n; i++ {
		maxRow := 0
		maxCol := 0
		for j := 0; j < n; j++ {
			maxRow = max(maxRow, grid[i][j])
			maxCol = max(maxCol, grid[j][i])
			xy += min(grid[i][j], 1)
		}
		yz += maxRow
		zx += maxCol
	}
	return xy + yz + zx
}
