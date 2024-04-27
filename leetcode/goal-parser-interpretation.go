func interpret(command string) string {
	output := []rune{}
	n := 0
	for _, c := range command {
		if c == 'G' || c == 'l' || c == '(' {
			output = append(output, c)
			n++
		} else if c == 'a' {
			output[n-1] = c
		} else if output[n-1] == '(' {
			output[n-1] = 'o'
		}
	}
	return string(output)
}
