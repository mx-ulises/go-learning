func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		return 'a' + c - 'A'
	}
	return c
}

func toUpper(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return 'A' + c - 'a'
	}
	return c
}

func getPermutations(s string, i int, current []byte, output []string) []string {
	if i == len(s) {
		return append(output, string(current))
	}
	c := s[i]
	i++
	current = append(current, toLower(c))
	output = getPermutations(s, i, current, output)
	if c < '0' || '9' < c {
		current[i-1] = toUpper(c)
		output = getPermutations(s, i, current, output)
	}
	i--
	current = current[:i]
	return output
}

func letterCasePermutation(s string) []string {
	return getPermutations(s, 0, []byte{}, []string{})
}
