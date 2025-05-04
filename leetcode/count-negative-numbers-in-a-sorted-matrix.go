func countNegatives(grid [][]int) int {
	count := 0
	i, m := 0, len(grid[0])
	for _, row := range grid {
		for i < m && row[m-i-1] < 0 {
			i++
		}
		count += i
	}
	return count
}
