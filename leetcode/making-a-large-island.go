func isValid(grid [][]int, x, y int) bool {
	n, m := len(grid), len(grid[0])
	return 0 <= x && x < n && 0 <= y && y < m
}

func getMergedIsland(islands map[int]int, i, j int, grid, neighbors [][]int) int {
	island := grid[i][j]
	if island != 0 {
		return islands[island]
	}
	neighborIslandVisited := map[int]bool{}
	islandSize := 1
	for _, offset := range neighbors {
		x, y := i+offset[0], j+offset[1]
		if isValid(grid, x, y) {
			island := grid[x][y]
			if neighborIslandVisited[island] == false {
				islandSize += islands[island]
				neighborIslandVisited[island] = true
			}
		}
	}
	return islandSize
}

func visit(i, j, island int, islands map[int]int, grid, neighbors [][]int) int {
	if isValid(grid, i, j) == false || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = island
	size := 1
	for _, offset := range neighbors {
		x, y := i+offset[0], j+offset[1]
		size += visit(x, y, island, islands, grid, neighbors)
	}
	return size
}

func getIslands(grid, neighbors [][]int) map[int]int {
	islands := map[int]int{0: 0}
	island := 2
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			islands[island] = 0
			size := visit(i, j, island, islands, grid, neighbors)
			if 0 < size {
				islands[island] = size
				island++
			}
		}
	}
	return islands
}

func largestIsland(grid [][]int) int {
	neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	islands := getIslands(grid, neighbors)
	largest := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			candidate := getMergedIsland(islands, i, j, grid, neighbors)
			largest = max(largest, candidate)
		}
	}
	return largest
}
