func longestSubarray(nums []int) int {
	maximal := 0
	maximalCount := 0
	current := 0
	currentCount := 0
	for _, num := range nums {
		if current == num {
			currentCount++
		} else {
			current = num
			currentCount = 1
		}
		if maximal < current {
			maximal = current
			maximalCount = currentCount
		}
		if maximal == current {
			maximalCount = max(maximalCount, currentCount)
		}
	}
	return maximalCount
}
