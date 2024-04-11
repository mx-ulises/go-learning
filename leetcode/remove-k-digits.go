func removeKdigits(num string, k int) string {
	stack, stackSize := []rune{}, 0
	for _, d := range num {
		for 0 < stackSize && 0 < k && d < stack[stackSize-1] {
			stackSize, k = stackSize-1, k-1
			stack = stack[:stackSize]
		}
		stack = append(stack, d)
		stackSize++
	}
	stackSize -= k
	stack = stack[:stackSize]
	for 0 < stackSize && stack[0] == '0' {
		stack, stackSize = stack[1:], stackSize-1
	}
	if stackSize == 0 {
		stack = []rune{'0'}
	}
	return string(stack)
}
