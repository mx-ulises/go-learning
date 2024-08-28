func IsLand(grid *[][]int, i, j int) bool {
	if i < 0 || len(*grid) <= i {
		return false
	}
	if j < 0 || len((*grid)[i]) <= j {
		return false
	}
	return (*grid)[i][j] == 1
}

func IsValidSubisland(grid1, grid2, moves *[][]int, i, j int) (bool, bool) {
	if IsLand(grid2, i, j) == false {
		// It is not SubIsland, but it is valid
		return false, true
	}
	// It isSubisland, remove land
	(*grid2)[i][j] = 0
	// Check if it is valid
	isValid := (*grid1)[i][j] == 1
	for _, move := range *moves {
		_, valid := IsValidSubisland(grid1, grid2, moves, i+move[0], j+move[1])
		isValid = isValid && valid
	}
	// Return if it is a subIsland that isValid
	return true, isValid
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	subIslands := 0
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			subIsland, isValid := IsValidSubisland(&grid1, &grid2, &moves, i, j)
			if subIsland && isValid {
				subIslands++
			}
		}
	}
	return subIslands
}
