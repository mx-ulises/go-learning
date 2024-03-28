func maxSubarrayLength(nums []int, k int) int {
	count := map[int]int{}
	left, right, n := 0, 0, len(nums)
	maximal := 0
	for right < n {
		// add last number to the subarray
		num := nums[right]
		count[num]++
		right++
		// if added num breach the threshold, remove elements in the left side
		for k < count[num] {
			count[nums[left]]--
			left++
		}
		// Check size of the current subarray and update maximal
		candidate := right - left
		if maximal < candidate {
			maximal = candidate
		}
	}
	return maximal
}
