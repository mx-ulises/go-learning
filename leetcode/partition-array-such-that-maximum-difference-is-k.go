package leetcode

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	count := 1
	minimal := nums[0]
	for i := 1; i < len(nums); i++ {
		if k < (nums[i] - minimal) {
			count++
			minimal = nums[i]
		}
	}
	return count
}
