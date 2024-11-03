func compare(s, goal string, j, n int) bool {
	for i := 0; i < n; i++ {
		if s[i] != goal[j] {
			return false
		}
		j++
		if j == n {
			j = 0
		}
	}
	return true
}

func rotateString(s string, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}
	for i := 0; i < n; i++ {
		if compare(s, goal, i, n) {
			return true
		}
	}
	return false
}
