package leetcode

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	totalApples := getTotalApples(apple)
	return getMinimalBoxCount(totalApples, capacity)
}

func getTotalApples(apple []int) int {
	totalApples := 0
	for _, apples := range apple {
		totalApples += apples
	}
	return totalApples
}

func getMinimalBoxCount(totalApples int, capacity []int) int {
	minimalBoxCount := 0
	sort.Ints(capacity)
	i := len(capacity) - 1
	for 0 < totalApples {
		minimalBoxCount++
		totalApples -= capacity[i]
		i--
	}
	return minimalBoxCount
}
