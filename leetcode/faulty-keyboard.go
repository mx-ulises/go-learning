func finalString(s string) string {
	output := []rune{}
	for _, c := range s {
		if c == 'i' {
			n := len(output)
			for i := 0; i < (n / 2); i++ {
				j := n - i - 1
				output[j], output[i] = output[i], output[j]
			}
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}
