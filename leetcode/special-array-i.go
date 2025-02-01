func isArraySpecial(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		parityLeft, parityRight := nums[i-1]&1, nums[i]&1
		if parityLeft == parityRight {
			return false
		}
	}
	return true
}
