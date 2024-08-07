func minOperations(nums []int) int {
	operations := 0
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		operations += max(0, current-nums[i]+1)
		current = max(nums[i], current+1)
	}
	return operations
}
