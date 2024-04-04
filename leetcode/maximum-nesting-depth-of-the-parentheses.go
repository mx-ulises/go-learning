func maxDepth(s string) int {
	maximal, stack := 0, 0
	for _, c := range s {
		if c == '(' {
			stack++
			if maximal < stack {
				maximal = stack
			}
		} else if c == ')' {
			stack--
		}
	}
	return maximal
}
