func divideArray(nums []int) bool {
	numCount := [501]int{}
	noPair := 0
	for _, num := range nums {
		numCount[num]++
		if numCount[num]&1 == 1 {
			noPair++
		} else {
			noPair--
		}
	}
	return noPair == 0
}
