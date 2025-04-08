func minimumOperations(nums []int) int {
	numCount, nonUniqueCount := map[int]int{}, 0
	for _, num := range nums {
		numCount[num]++
		if numCount[num] == 2 {
			nonUniqueCount++
		}
	}
	i := 0
	for 0 < nonUniqueCount {
		num := nums[i]
		if numCount[num] == 2 {
			nonUniqueCount--
		}
		numCount[num]--
		i++
	}
	i += 2
	return i / 3
}
