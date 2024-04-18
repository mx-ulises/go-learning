func IsWater(grid *[][]int, i, j, n, m int) bool {
	if i < 0 || i == n || j < 0 || j == m {
		return true
	}
	return (*grid)[i][j] == 0
}

func GetLandPerimeter(grid *[][]int, neighbors *[][]int, i, j, n, m int) int {
	perimeter := 0
	if (*grid)[i][j] == 1 {
		for _, neighbor := range *neighbors {
			if IsWater(grid, i+neighbor[0], j+neighbor[1], n, m) {
				perimeter++
			}
		}
	}
	return perimeter
}

func islandPerimeter(grid [][]int) int {
	neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	perimeter := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			perimeter += GetLandPerimeter(&grid, &neighbors, i, j, n, m)
		}
	}
	return perimeter
}
