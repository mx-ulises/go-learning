package leetcode

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	i, minimumCost := len(cost)-1, 0
	for 2 <= i {
		minimumCost += cost[i] + cost[i-1]
		i -= 3
	}
	for 0 <= i {
		minimumCost += cost[i]
		i--
	}
	return minimumCost
}
