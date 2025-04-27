func distinctDifferenceArray(nums []int) []int {
	uniqueLeft := [51]int{}
	uniqueLeftCount := 0
	uniqueRight := [51]int{}
	uniqueRightCount := 0
	diff := make([]int, len(nums))
	for _, num := range nums {
		if uniqueRight[num] == 0 {
			uniqueRightCount++
		}
		uniqueRight[num]++
	}
	for i, num := range nums {
		if uniqueLeft[num] == 0 {
			uniqueLeftCount++
		}
		uniqueLeft[num]++
		uniqueRight[num]--
		if uniqueRight[num] == 0 {
			uniqueRightCount--
		}
		diff[i] = uniqueLeftCount - uniqueRightCount
	}
	return diff
}
