func EvaluateCount(l, r, w, dir int) bool {
	if dir == 1 {
		return (l + w) < r
	}
	return (r + w) < l
}

func ValidateString(s *string, i, end, dir int) bool {
	count := map[byte]int{}
	for i != end {
		count[(*s)[i]]++
		if EvaluateCount(count['('], count[')'], count['*'], dir) {
			return false
		}
		i += dir
	}
	return true
}

func checkValidString(s string) bool {
	n := len(s)
	return ValidateString(&s, 0, n, 1) && ValidateString(&s, n-1, -1, -1)
}
