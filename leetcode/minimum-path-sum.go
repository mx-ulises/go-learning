func GetMinPath(grid *[][]int, distance *[][]int, x int, y int) int {
	if x == len(*grid) {
		return 1000000
	}
	if y == len((*grid)[x]) {
		return 1000000
	}
	if (*distance)[x][y] == -1 {
		candidate1 := (*grid)[x][y] + GetMinPath(grid, distance, x+1, y)
		candidate2 := (*grid)[x][y] + GetMinPath(grid, distance, x, y+1)
		if candidate1 < candidate2 {
			(*distance)[x][y] = candidate1
		} else {
			(*distance)[x][y] = candidate2
		}
	}
	return (*distance)[x][y]
}

func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	distance := [][]int{}
	for i := 0; i < n; i++ {
		distance = append(distance, []int{})
		for j := 0; j < m; j++ {
			distance[i] = append(distance[i], -1)
		}
	}
	distance[n-1][m-1] = grid[n-1][m-1]
	GetMinPath(&grid, &distance, 0, 0)
	return distance[0][0]
}
