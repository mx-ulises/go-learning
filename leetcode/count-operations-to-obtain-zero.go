func countOperations(num1 int, num2 int) int {
	operations := 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			operations += num1 / num2
			num1 %= num2
		} else {
			operations += num2 / num1
			num2 %= num1
		}
	}
	return operations
}
