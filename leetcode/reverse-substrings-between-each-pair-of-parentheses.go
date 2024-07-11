func reverseParentheses(s string) string {
	output := []rune{}
	for _, c := range s {
		if c != ')' {
			output = append(output, c)
		} else {
			buffer := []rune{}
			n := len(output) - 1
			for output[n] != '(' {
				buffer = append(buffer, output[n])
				output = output[:n]
				n--
			}
			output = output[:n]
			for _, c := range buffer {
				output = append(output, c)
			}
		}
	}
	return string(output)
}
