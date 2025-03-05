func transformArray(nums []int) []int {
	for i, num := range nums {
		nums[i] = num & 1
	}
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l] == 0 {
			l++
		}
		for l < r && nums[r] == 1 {
			r--
		}
		aux := nums[l]
		nums[l] = nums[r]
		nums[r] = aux
	}
	return nums
}
