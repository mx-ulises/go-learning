func removeStars(s string) string {
	stack, n := []rune{}, 0
	for _, c := range s {
		if c == '*' {
			n--
			stack = stack[:n]
		} else {
			n++
			stack = append(stack, c)
		}
	}
	return string(stack)
}
