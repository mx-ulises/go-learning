func minCosts(cost []int) []int {
	n := len(cost)
	answer := make([]int, n)
	answer[0] = cost[0]
	for i := 1; i < n; i++ {
		answer[i] = min(answer[i-1], cost[i])
	}
	return answer
}
