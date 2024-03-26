func minPairSum(nums []int) int {
	sort.Ints(nums)
	maxSum := 0
	start, end := 0, len(nums)-1
	for start < end {
		candidate := nums[start] + nums[end]
		if maxSum < candidate {
			maxSum = candidate
		}
		start++
		end--
	}
	return maxSum
}
