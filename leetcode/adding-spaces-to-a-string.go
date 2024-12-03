func addSpaces(s string, spaces []int) string {
	output := make([]string, len(spaces)+1)
	start := 0
	i := 0
	for _, space := range spaces {
		output[i] = s[start:space]
		start = space
		i++
	}
	output[i] = s[start:]
	return strings.Join(output, " ")
}
