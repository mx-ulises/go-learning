func xorOperation(n int, start int) int {
	total := 0
	for i := 0; i < n; i++ {
		total ^= start + 2*i
	}
	return total
}
