func differenceOfSums(n int, m int) int {
	diff := n * (n + 1) / 2
	num2 := 0
	for x := m; x <= n; x += m {
		num2 += x
	}
	return diff - num2 - num2
}
