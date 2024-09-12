func largestAltitude(gain []int) int {
	current := 0
	maximal := 0
	for _, h := range gain {
		current += h
		maximal = max(maximal, current)
	}
	return maximal
}
