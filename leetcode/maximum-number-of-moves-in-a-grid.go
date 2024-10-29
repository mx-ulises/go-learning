func IsValidPosition(grid *[][]int, i, j, prev int) bool {
	n, m := len(*grid), len((*grid)[0])
	if i < 0 || i == n {
		return false
	}
	if j < 0 || j == m {
		return false
	}
	return prev < (*grid)[i][j]
}

func GetMoves(grid *[][]int, visited *map[int]int, validMoves *[][2]int, i, j int) int {
	code := i*1000 + j
	if _, ok := (*visited)[code]; ok {
		return (*visited)[code]
	}
	maximal := 0
	for _, move := range *validMoves {
		iNext, jNext := i+move[0], j+move[1]
		if IsValidPosition(grid, iNext, jNext, (*grid)[i][j]) {
			maximal = max(maximal, 1+GetMoves(grid, visited, validMoves, iNext, jNext))
		}
	}
	(*visited)[code] = maximal
	return (*visited)[code]
}

func maxMoves(grid [][]int) int {
	validMoves := [][2]int{{-1, 1}, {0, 1}, {1, 1}}
	n := len(grid)
	maximal := 0
	visited := map[int]int{}
	for i := 0; i < n; i++ {
		candidate := GetMoves(&grid, &visited, &validMoves, i, 0)
		maximal = max(maximal, candidate)
	}
	return maximal
}
