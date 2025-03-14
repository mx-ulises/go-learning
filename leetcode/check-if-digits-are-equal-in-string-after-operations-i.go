func hasSameDigits(s string) bool {
	digits := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = s[i] - '0'
	}
	for 2 < len(digits) {
		d, i := digits[len(digits)-1], len(digits)-1
		digits = digits[:i]
		for 0 < i {
			j := i - 1
			aux := digits[j]
			digits[j] += d
			if 10 <= digits[j] {
				digits[j] -= 10
			}
			d, i = aux, j
		}
	}
	return digits[0] == digits[1]
}
