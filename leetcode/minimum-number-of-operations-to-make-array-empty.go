func minOperations(nums []int) int {
	numCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numCount[nums[i]] += 1
	}
	operationCount := 0
	for _, count := range numCount {
		if count == 1 {
			return -1
		}
		operationCount += count / 3
		if (count % 3) != 0 {
			operationCount += 1
		}
	}
	return operationCount
}
