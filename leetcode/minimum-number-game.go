func numberGame(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}
