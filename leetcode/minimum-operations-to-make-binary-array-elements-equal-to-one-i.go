func minOperations(nums []int) int {
	operations := 0
	n := len(nums) - 2
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			operations++
		}
	}
	for i := n; i < len(nums); i++ {
		if nums[i] != 1 {
			return -1
		}
	}
	return operations
}
