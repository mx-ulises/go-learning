package leetcode

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	difference := 100000
	limit := len(nums) - k + 1
	for i := 0; i < limit; i++ {
		j := i + k - 1
		difference = min(difference, nums[j]-nums[i])
	}
	return difference
}
