package leetcode

func getDigit(c int) byte {
	if 0 <= c && c <= 9 {
		return '0' + byte(c)
	}
	return 'A' + byte(c-10)
}

func reverse(s []byte) []byte {
	n := len(s) / 2
	for i := 0; i < n; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func concatHex36(n int) string {
	n2 := n * n
	n3 := n * n2
	output, aux := []byte{}, []byte{}
	for 0 < n2 {
		output = append(output, getDigit(n2%16))
		n2 /= 16
	}
	for 0 < n3 {
		aux = append(aux, getDigit(n3%36))
		n3 /= 36
	}
	output = append(reverse(output), reverse(aux)...)
	return string(output)
}
