func canJump(nums []int) bool {
	maximal := 0
	for i := 0; i < len(nums); i++ {
		if maximal < i {
			return false
		}
		candidate := i + nums[i]
		if maximal < candidate {
			maximal = candidate
		}
	}
	return true
}
