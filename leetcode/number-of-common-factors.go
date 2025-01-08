func commonFactors(a int, b int) int {
	minimal := min(a, b)
	count := 0
	for i := 1; i <= minimal; i++ {
		if (a%i) == 0 && (b%i) == 0 {
			count++
		}
	}
	return count
}
