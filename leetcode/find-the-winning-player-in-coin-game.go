package leetcode

func winningPlayer(x int, y int) string {
	gameTurns := min(x, y/4)
	if isOdd(gameTurns) {
		return "Alice"
	}
	return "Bob"
}

func isOdd(x int) bool {
	return x&1 == 1
}
