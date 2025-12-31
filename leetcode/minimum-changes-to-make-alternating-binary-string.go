package leetcode

func minOperations(s string) int {
	operationsStartingWithZero := getOperationsStartingWith(s, '0')
	operationsStartingWithOne := getOperationsStartingWith(s, '1')
	return min(operationsStartingWithZero, operationsStartingWithOne)
}

func getOperationsStartingWith(s string, c byte) int {
	operations := 0
	for i := 0; i < len(s); i++ {
		if s[i] != c {
			operations++
		}
		c = flipCharacter(c)
	}
	return operations
}

func flipCharacter(c byte) byte {
	if c == '0' {
		return '1'
	}
	return '0'
}
