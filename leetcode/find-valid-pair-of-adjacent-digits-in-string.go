package leetcode

func findValidPair(s string) string {
	digitCount := getDigitCount(s)
	prev := s[0]
	for i := 1; i < len(s); i++ {
		if isValidPair(digitCount, prev, s[i]) {
			return s[i-1 : i+1]
		}
		prev = s[i]
	}
	return ""
}

func getDigitCount(s string) [10]int {
	digitCount := [10]int{}
	for i := 0; i < len(s); i++ {
		digit := getDigit(s[i])
		digitCount[digit]++
	}
	return digitCount
}

func getDigit(c byte) int {
	return int(c - '0')
}

func isValidPair(digitCount [10]int, prev byte, cur byte) bool {
	prevValid, curValid := isValidDigit(digitCount, prev), isValidDigit(digitCount, cur)
	return prev != cur && prevValid && curValid
}

func isValidDigit(digitCount [10]int, c byte) bool {
	digit := getDigit(c)
	return digitCount[digit] == digit
}
