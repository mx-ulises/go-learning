func finalValueAfterOperations(operations []string) int {
	n, value := len(operations), 0
	for i := 0; i < n; i++ {
		if operations[i][1] == '-' {
			value--
		} else {
			value++
		}
	}
	return value
}
