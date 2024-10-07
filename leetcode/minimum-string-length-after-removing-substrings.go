func minLength(s string) int {
	stack := []rune{}
	n := 0
	for _, c := range s {
		if 0 < n && stack[n-1] == 'A' && c == 'B' {
			n--
			stack = stack[:n]
		} else if 0 < n && stack[n-1] == 'C' && c == 'D' {
			n--
			stack = stack[:n]
		} else {
			stack = append(stack, c)
			n++
		}
	}
	return n
}
