func diStringMatch(s string) []int {
	l, h := 0, len(s)
	output := []int{}
	for _, c := range s {
		if c == 'I' {
			output = append(output, l)
			l++
		} else {
			output = append(output, h)
			h--
		}
	}
	output = append(output, l)
	return output
}
