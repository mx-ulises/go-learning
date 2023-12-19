func maxProductDifference(nums []int) int {
	a := nums[0]
	b := nums[1]
	if a < b {
		b = nums[0]
		a = nums[1]
	}
	maximals := make([]int, 0)
	minimals := make([]int, 0)
	maximals = append(maximals, a)
	maximals = append(maximals, b)
	minimals = append(minimals, b)
	minimals = append(minimals, a)
	for i := 2; i < len(nums); i++ {
		c := nums[i]
		if maximals[0] <= c {
			maximals[1] = maximals[0]
			maximals[0] = c
		} else if maximals[1] < c {
			maximals[1] = c
		}
		if c <= minimals[0] {
			minimals[1] = minimals[0]
			minimals[0] = c
		} else if c < minimals[1] {
			minimals[1] = c
		}
	}
	return (maximals[0] * maximals[1]) - (minimals[0] * minimals[1])
}
