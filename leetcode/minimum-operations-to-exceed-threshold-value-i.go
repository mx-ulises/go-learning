func minOperations(nums []int, k int) int {
	operationsCount := 0
	for _, num := range nums {
		if num < k {
			operationsCount += 1
		}
	}
	return operationsCount
}
