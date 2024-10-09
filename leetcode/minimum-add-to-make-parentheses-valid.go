func minAddToMakeValid(s string) int {
	stack := []rune{}
	moves := 0
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			moves++
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	moves += len(stack)
	return moves
}
