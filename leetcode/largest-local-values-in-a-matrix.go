func GetLargestLocal(grid *[][]int, x, y int) int {
	largest := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			largest = max(largest, (*grid)[x+i][y+j])
		}
	}
	return largest
}

func largestLocal(grid [][]int) [][]int {
	n := len(grid) - 2
	output := make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, n)
		for j := 0; j < n; j++ {
			output[i][j] = GetLargestLocal(&grid, i, j)
		}
	}
	return output
}
