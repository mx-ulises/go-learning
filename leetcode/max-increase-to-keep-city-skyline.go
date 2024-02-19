func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	xMaximals := make([]int, n)
	yMaximals := make([]int, n)
	output := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			xMaximals[i] = max(xMaximals[i], grid[i][j])
			yMaximals[j] = max(yMaximals[j], grid[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			minimal := min(xMaximals[i], yMaximals[j])
			output += max(0, minimal-grid[i][j])
		}
	}
	return output
}
