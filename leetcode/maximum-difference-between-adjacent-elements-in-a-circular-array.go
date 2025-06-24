func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	maxDiff := abs(nums[n-1] - nums[0])
	for i := 1; i < n; i++ {
		maxDiff = max(maxDiff, abs(nums[i-1]-nums[i]))
	}
	return maxDiff
}
