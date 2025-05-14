func digitCount(num string) bool {
	digitCount := [10]int{}
	for i, digit := range num {
		d := int(digit - '0')
		digitCount[d]++
		digitCount[i] -= d
	}
	for _, d := range digitCount {
		if d != 0 {
			return false
		}
	}
	return true
}
