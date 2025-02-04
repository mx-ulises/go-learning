func maxAscendingSum(nums []int) int {
	maximal, current := 0, 0
	prev := 0
	for _, num := range nums {
		if prev < num {
			current += num
		} else {
			current = num
		}
		maximal = max(maximal, current)
		prev = num
	}
	return maximal
}
