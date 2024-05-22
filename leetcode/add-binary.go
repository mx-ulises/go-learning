func reverse(output *[]byte) {
	n := len(*output)
	m := n / 2
	for i := 0; i < m; i++ {
		j := n - i - 1
		aux := (*output)[i]
		(*output)[i] = (*output)[j]
		(*output)[j] = aux
	}
}

func addBinary(a string, b string) string {
	output := []byte{}
	i, j, carry := len(a)-1, len(b)-1, 0
	mapToDigit := map[byte]int{'1': 1, '0': 0}
	mapToChar := []byte{'0', '1'}
	for 0 <= i || 0 <= j || carry == 1 {
		if 0 <= i {
			carry += mapToDigit[a[i]]
			i--
		}
		if 0 <= j {
			carry += mapToDigit[b[j]]
			j--
		}
		value := mapToChar[carry&1]
		carry /= 2
		output = append(output, value)
	}
	reverse(&output)
	return string(output)
}
