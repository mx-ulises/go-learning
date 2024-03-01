func onesMinusZeros(grid [][]int) [][]int {
	diffRows := make([]int, len(grid))
	diffCols := make([]int, len(grid[0]))
	diffMap := []int{-1, 1}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			diffRows[i] += diffMap[grid[i][j]]
			diffCols[j] += diffMap[grid[i][j]]
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = diffRows[i] + diffCols[j]
		}
	}
	return grid
}
