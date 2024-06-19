func decode(encoded []int, first int) []int {
	original := make([]int, len(encoded)+1)
	original[0] = first
	for i, x := range encoded {
		original[i+1] = x ^ original[i]
	}
	return original
}
