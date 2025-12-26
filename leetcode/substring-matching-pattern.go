package leetcode

func hasMatch(s string, p string) bool {
	leftPattern := getLeftPattern(p)
	rightPattern := getRightPattern(p, len(leftPattern)+1)
	leftOk, leftEndIndex := testLeft(s, leftPattern)
	rightOk, _ := testRight(s, rightPattern, leftEndIndex)
	return leftOk && rightOk
}

func getLeftPattern(p string) string {
	leftPattern, i := []byte{}, 0
	for p[i] != '*' {
		leftPattern = append(leftPattern, p[i])
		i++
	}
	return string(leftPattern)
}

func getRightPattern(p string, i int) string {
	rightPattern := []byte{}
	for i < len(p) {
		rightPattern = append(rightPattern, p[i])
		i++
	}
	return string(rightPattern)
}

func testLeft(s string, p string) (bool, int) {
	if len(p) == 0 {
		return true, 0
	}
	for i := 0; i < len(s); i++ {
		if testPattern(s, i, p) {
			return true, i + len(p)
		}
	}
	return false, len(s)
}

func testRight(s string, p string, start int) (bool, int) {
	if len(p) == 0 {
		return true, start
	}
	for i := start; i < len(s); i++ {
		if testPattern(s, i, p) {
			return true, i
		}
	}
	return false, len(s)
}

func testPattern(s string, j int, p string) bool {
	for i := 0; i < len(p); i++ {
		if j == len(s) || s[j] != p[i] {
			return false
		}
		j++
	}
	return true
}
