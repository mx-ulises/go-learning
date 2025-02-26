func maxAbsoluteSum(nums []int) int {
	minimal, maximal := 0, 0
	current := 0
	for _, num := range nums {
		current += num
		minimal = min(minimal, current)
		maximal = max(maximal, current)
	}
	return maximal - minimal
}
