func toLowerCase(s string) string {
	output := []rune{}
	for _, c := range s {
		if 'A' <= c && c <= 'Z' {
			output = append(output, c+32)
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}
