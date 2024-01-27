func isValid(s string) bool {
	parenthesisMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := []rune{}
	for _, c := range s {
		d, ok := parenthesisMap[c]
		if ok {
			n := len(stack) - 1
			if n == -1 || stack[n] != d {
				return false
			}
			stack = stack[:n]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
