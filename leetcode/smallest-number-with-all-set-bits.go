func smallestNumber(n int) int {
	x := 1
	for x < n {
		x <<= 1
		x |= 1
	}
	return x
}
