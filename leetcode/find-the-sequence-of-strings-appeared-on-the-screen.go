func stringSequence(target string) []string {
	s := []byte{'a'}
	i := 0
	output := []string{}
	for i < len(target) {
		output = append(output, string(s))
		if s[i] < target[i] {
			s[i]++
		} else {
			s = append(s, 'a')
			i++
		}
	}
	return output
}
