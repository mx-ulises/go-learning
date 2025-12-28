package leetcode

func canMakeSquare(grid [][]byte) bool {
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid); j++ {
			if canCreateSquare(grid, i, j) {
				return true
			}
		}
	}
	return false
}

func canCreateSquare(grid [][]byte, i int, j int) bool {
	return countWhiteCells(grid, i, j) != 2
}

func countWhiteCells(grid [][]byte, i int, j int) int {
	whiteSquares := 0
	for offsetX := 0; offsetX < 2; offsetX++ {
		for offsetY := 0; offsetY < 2; offsetY++ {
			whiteSquares += returnOneIfWhite(grid[i-offsetX][j-offsetY])
		}
	}
	return whiteSquares
}

func returnOneIfWhite(cell byte) int {
	if cell == 'W' {
		return 1
	}
	return 0
}
