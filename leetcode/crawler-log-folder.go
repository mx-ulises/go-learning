func minOperations(logs []string) int {
	operationEffect := map[string]int{"../": -2, "./": -1}
	minOperations := 0
	for _, operation := range logs {
		minOperations = max(0, minOperations+operationEffect[operation]+1)
	}
	return minOperations
}
