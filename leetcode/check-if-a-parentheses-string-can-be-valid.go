func testString(s, locked string, start, end, step int, addChar byte) bool {
	balance, wildcards := 0, 0
	for start != end {
		if locked[start] == '0' {
			wildcards++
		} else if s[start] == addChar {
			balance++
		} else {
			balance--
		}
		if (balance + wildcards) < 0 {
			return false
		}
		start += step
	}
	return balance <= wildcards
}

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}
	leftToRightTest := testString(s, locked, 0, n, 1, '(')
	rightToLeftTest := testString(s, locked, n-1, -1, -1, ')')
	return leftToRightTest && rightToLeftTest
}
