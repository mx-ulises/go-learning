func maxProduct(nums []int) int {
	a, b := max(nums[0], nums[1]), min(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		if a < nums[i] {
			b = a
			a = nums[i]
		} else if b < nums[i] {
			b = nums[i]
		}
	}
	return (a - 1) * (b - 1)
}
