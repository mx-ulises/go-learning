func repeatedNTimes(nums []int) int {
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
