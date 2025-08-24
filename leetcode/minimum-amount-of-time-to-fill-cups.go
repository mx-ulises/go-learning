package leetcode

import "sort"

func fillCups(amount []int) int {
	total := 0
	sort.Ints(amount)
	for 0 < amount[2] {
		total++
		amount[1], amount[2] = amount[1]-1, amount[2]-1
		sort.Ints(amount)
	}
	return total
}
