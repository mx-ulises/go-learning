func findChampion(n int, edges [][]int) int {
	losers := make([]bool, n)
	for _, edge := range edges {
		losers[edge[1]] = true
	}
	champion, championCount := -1, 0
	for i, loser := range losers {
		if loser == false {
			champion = i
			championCount++
			if 1 < championCount {
				champion = -1
				break
			}
		}
	}
	return champion
}
