func makeGood(s string) string {
	stack, n := make([]rune, 0), 0
	diff := 'a' - 'A'
	for _, c := range s {
		if 0 < n && (stack[n-1]+diff == c || c+diff == stack[n-1]) {
			stack = stack[:n-1]
			n--
		} else {
			stack = append(stack, c)
			n++
		}
	}
	return string(stack)
}
