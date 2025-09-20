package leetcode

import "sort"

func maxKDistinct(nums []int, k int) []int {
	output := []int{}
	sort.Ints(nums)
	i, j := len(nums)-1, 0
	output = append(output, nums[i])
	for 0 <= i && len(output) < k {
		if output[j] != nums[i] {
			output = append(output, nums[i])
			j++
		}
		i--
	}
	return output
}
