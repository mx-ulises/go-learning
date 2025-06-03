func smallestRangeI(nums []int, k int) int {
	maximal, minimal := 0, 10000
	for _, num := range nums {
		maximal = max(maximal, num)
		minimal = min(minimal, num)
	}
	return max(maximal-minimal-2*k, 0)
}
