func minimumOperations(nums []int) int {
	operationCount := 0
	for _, num := range nums {
		if (num % 3) != 0 {
			operationCount++
		}
	}
	return operationCount
}
