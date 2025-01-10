func minimumChairs(s string) int {
	operation := map[byte]int{'E': 1, 'L': -1}
	chairs, minimum := 0, 0
	for i := 0; i < len(s); i++ {
		chairs += operation[s[i]]
		minimum = max(minimum, chairs)
	}
	return minimum
}
