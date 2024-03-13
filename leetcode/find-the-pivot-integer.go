func pivotInteger(n int) int {
	right := n * (n + 1) / 2
	left := 0
	for i := 1; i <= n; i++ {
		right -= i
		if right == left {
			return i
		}
		left += i
	}
	return -1
}
