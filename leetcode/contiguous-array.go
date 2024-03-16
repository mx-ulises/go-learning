func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMaxLength(nums []int) int {
	maxLength := 0
	sumPositionMap := map[int]int{0: 0}
	sumMap := map[int]int{0: -1, 1: 1}
	sum := 0
	for i, x := range nums {
		sum += sumMap[x]
		position, ok := sumPositionMap[sum]
		if ok {
			maxLength = max(maxLength, i+1-position)
		} else {
			sumPositionMap[sum] = i + 1
		}
	}
	return maxLength
}
