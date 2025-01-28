func totalMoney(n int) int {
	total, extra := 0, 0
	for 7 <= n {
		total += 28 + extra*7
		n -= 7
		extra++
	}
	for i := 1; i < (n + 1); i++ {
		total += i + extra
	}
	return total
}
