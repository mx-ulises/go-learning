func AddReversedWord(output *[]rune, buffer *[]rune) {
	n := len(*buffer)
	for i := 0; i < n; i++ {
		*output = append(*output, (*buffer)[n-i-1])
	}
}

func reverseWords(s string) string {
	buffer, output := []rune{}, []rune{}
	for _, c := range s {
		if c == ' ' {
			AddReversedWord(&output, &buffer)
			output = append(output, ' ')
			buffer = []rune{}
		} else {
			buffer = append(buffer, c)
		}
	}
	AddReversedWord(&output, &buffer)
	return string(output)
}
