func removeOuterParentheses(s string) string {
	output := []byte{}
	stackSize := 0
	for _, c := range s {
		if c == '(' {
			if stackSize != 0 {
				output = append(output, '(')
			}
			stackSize++
		}
		if c == ')' {
			stackSize--
			if stackSize != 0 {
				output = append(output, ')')
			}
		}
	}
	return string(output)
}
