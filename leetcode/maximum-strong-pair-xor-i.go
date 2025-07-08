package leetcode

import "sort"

func getOrderedSet(nums []int) []int {
	sort.Ints(nums)
	end := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[end] = nums[i]
			end++
		}
	}
	return nums[:end]
}

func maximumStrongPairXor(nums []int) int {
	nums = getOrderedSet(nums)
	maximal, left, right := 0, 0, 0
	for right < len(nums) {
		for nums[left] < (nums[right] - nums[left]) {
			left++
		}
		for i := left; i <= right; i++ {
			maximal = max(maximal, nums[right]^nums[i])
		}
		right++
	}
	return maximal
}
