func maximumOddBinaryNumber(s string) string {
	ones, zeros := []rune{}, []rune{}
	for _, c := range s {
		if c == '0' {
			zeros = append(zeros, '0')
		} else {
			ones = append(ones, '1')
		}
	}
	ones = append(ones, zeros...)
	ones = append(ones, '1')
	return string(ones[1:])
}
