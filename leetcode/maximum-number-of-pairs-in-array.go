func numberOfPairs(nums []int) []int {
	numCount := map[int]int{}
	output := []int{0, 0}
	for _, num := range nums {
		numCount[num]++
		if numCount[num] == 1 {
			output[1]++
		} else {
			output[1]--
			output[0]++
			numCount[num] = 0
		}
	}
	return output
}
