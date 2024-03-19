func reverseString(s []byte) {
	n := len(s)
	m := n / 2
	for i := 0; i < m; i++ {
		j := n - i - 1
		aux := s[i]
		s[i] = s[j]
		s[j] = aux
	}
}
