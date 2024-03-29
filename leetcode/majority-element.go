func majorityElement(nums []int) int {
	current, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == current {
			count++
		} else {
			count--
		}
		if count == -1 {
			current = nums[i]
			count = 1
		}
	}
	return current
}
