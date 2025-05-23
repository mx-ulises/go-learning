func findNonMinOrMax(nums []int) int {
	minimal, maximal := 101, 0
	for _, num := range nums {
		minimal = min(minimal, num)
		maximal = max(maximal, num)
	}
	for _, num := range nums {
		if num != minimal && num != maximal {
			return num
		}
	}
	return -1
}
