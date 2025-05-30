func minimizedStringLength(s string) int {
	visited := [26]bool{}
	minimized := len(s)
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		if visited[c] {
			minimized--
		} else {
			visited[c] = true
		}
	}
	return minimized
}
