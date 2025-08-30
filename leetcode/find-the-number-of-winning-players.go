package leetcode

func winningPlayerCount(n int, pick [][]int) int {
	playerBalls := make([][11]int, n)
	winners := make([]bool, n)
	for _, p := range pick {
		x, y := p[0], p[1]
		playerBalls[x][y]++
		if x < playerBalls[x][y] {
			winners[x] = true
		}
	}
	winnderCount := 0
	for _, won := range winners {
		if won {
			winnderCount++
		}
	}
	return winnderCount
}
