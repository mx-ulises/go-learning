package leetcode

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	end := y + k
	for y < end {
		grid = reverseColumn(grid, x, y, k)
		y++
	}
	return grid
}

func reverseColumn(grid [][]int, x int, y int, k int) [][]int {
	top, bottom := x, x+k-1
	for top < bottom {
		grid = swapCellsInColumnY(grid, top, bottom, y)
		top++
		bottom--
	}
	return grid
}

func swapCellsInColumnY(grid [][]int, top int, bottom int, y int) [][]int {
	aux := grid[top][y]
	grid[top][y] = grid[bottom][y]
	grid[bottom][y] = aux
	return grid
}
