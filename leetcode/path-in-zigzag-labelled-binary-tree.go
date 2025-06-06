func reverse(path []int) []int {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func reverseLabel(label int) int {
	start, end, dir := 0, 1, 1
	for end <= label {
		start, end = end, end+end
		dir ^= 1
	}
	if dir == 1 {
		label = end - label + start - 1
	}
	return label
}

func pathInZigZagTree(label int) []int {
	path := []int{reverseLabel(label)}
	i := 0
	for path[i] != 1 {
		path = append(path, path[i]/2)
		path[i] = reverseLabel(path[i])
		i += 1
	}
	return reverse(path)
}
