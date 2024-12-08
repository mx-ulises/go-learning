func coinChange(coins []int, amount int) int {
	q := [][]int{{0, 0}}
	visited := map[int]bool{}
	for 0 < len(q) {
		candidate, coinCount := q[0][0], q[0][1]
		q = q[1:]
		if candidate == amount {
			return coinCount
		}
		if visited[candidate] || amount < candidate {
			continue
		}
		coinCount++
		visited[candidate] = true
		for _, coin := range coins {
			q = append(q, []int{candidate + coin, coinCount})
		}
	}
	return -1
}
