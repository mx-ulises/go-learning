package leetcode

import "sort"

func smallestAbsent(nums []int) int {
	sort.Ints(nums)
	total, n := 0, len(nums)
	for _, num := range nums {
		total += num
	}
	candidate := max(total/n+1, 1)
	for _, num := range nums {
		if num == candidate {
			candidate++
		}
		if candidate < num {
			return candidate
		}
	}
	return candidate
}
