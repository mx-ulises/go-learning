func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]&1 == 0 {
			left++
		}
		for left < right && nums[right]&1 == 1 {
			right--
		}
		aux := nums[left]
		nums[left] = nums[right]
		nums[right] = aux
		left++
		right--
	}
	return nums
}
