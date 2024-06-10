func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	for i, h := range heights {
		expected[i] = h
	}
	sort.Ints(expected)
	count := 0
	for i, h := range heights {
		if h != expected[i] {
			count++
		}
	}
	return count
}
