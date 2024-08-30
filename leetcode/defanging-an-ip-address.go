func defangIPaddr(address string) string {
	output := make([]rune, len(address)+6)
	i := 0
	for _, c := range address {
		if c == '.' {
			output[i] = '['
			output[i+1] = '.'
			output[i+2] = ']'
			i += 3
		} else {
			output[i] = c
			i++
		}
	}
	return string(output)
}
