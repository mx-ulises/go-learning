func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	output := []int{}
	leftSum, r, rightSum := 0, len(nums)-1, 0
	for _, num := range nums {
		leftSum += num
	}
	for rightSum <= leftSum {
		output = append(output, nums[r])
		rightSum += nums[r]
		leftSum -= nums[r]
		r--
	}
	return output
}
