type NeighborSum struct {
	Grid     [][]int
	Adjacent map[int]int
	Diagonal map[int]int
}

func getValue(grid [][]int, x int, y int) int {
	if x < 0 || x == len(grid) || y < 0 || y == len(grid[0]) {
		return 0
	}
	return grid[x][y]
}

func Constructor(grid [][]int) NeighborSum {
	adjacentPositions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	diagonalPositions := [][]int{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	adjacent := map[int]int{}
	diagonal := map[int]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			num := grid[i][j]
			for _, pos := range adjacentPositions {
				adjacent[num] += getValue(grid, i+pos[0], j+pos[1])
			}
			for _, pos := range diagonalPositions {
				diagonal[num] += getValue(grid, i+pos[0], j+pos[1])
			}
		}
	}
	return NeighborSum{Adjacent: adjacent, Diagonal: diagonal}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	return this.Adjacent[value]
}

func (this *NeighborSum) DiagonalSum(value int) int {
	return this.Diagonal[value]
}

/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
