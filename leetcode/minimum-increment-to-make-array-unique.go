func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	current := 0
	increments := 0
	for _, num := range nums {
		current = max(current, num)
		increments += current - num
		current++
	}
	return increments
}
