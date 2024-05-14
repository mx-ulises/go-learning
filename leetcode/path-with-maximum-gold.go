func IsValid(grid *[][]int, i int, j int) bool {
	if i < 0 || len(*grid) <= i || j < 0 || len((*grid)[i]) <= j {
		return false
	}
	return (*grid)[i][j] != 0
}

func GoldPath(grid *[][]int, i int, j int, maximumGold int) int {
	if IsValid(grid, i, j) == false {
		return 0
	}
	value := (*grid)[i][j]
	(*grid)[i][j] = 0
	maximumGold += value
	current := maximumGold
	maximumGold = max(maximumGold, GoldPath(grid, i+1, j, current))
	maximumGold = max(maximumGold, GoldPath(grid, i-1, j, current))
	maximumGold = max(maximumGold, GoldPath(grid, i, j+1, current))
	maximumGold = max(maximumGold, GoldPath(grid, i, j-1, current))
	(*grid)[i][j] = value
	return maximumGold
}

func getMaximumGold(grid [][]int) int {
	maximumGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			maximumGold = max(maximumGold, GoldPath(&grid, i, j, 0))
		}
	}
	return maximumGold
}
