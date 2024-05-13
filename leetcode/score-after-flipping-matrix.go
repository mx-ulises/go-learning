func matrixScore(grid [][]int) int {
	score := 0
	for j := 0; j < len(grid[0]); j++ {
		totalZeros, totalOnes := 0, 0
		for i := 0; i < len(grid); i++ {
			if (grid[i][j] ^ grid[i][0]) == 0 {
				totalOnes++
			} else {
				totalZeros++
			}
		}
		score <<= 1
		score += max(totalZeros, totalOnes)
	}
	return score
}
