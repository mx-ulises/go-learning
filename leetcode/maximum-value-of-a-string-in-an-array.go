package leetcode

func getValue(s string) int {
	value := 0
	for _, c := range s {
		if '0' <= c && c <= '9' {
			value = 10*value + int(c-'0')
		} else {
			return len(s)
		}
	}
	return value
}

func maximumValue(strs []string) int {
	maximal := 0
	for _, s := range strs {
		maximal = max(maximal, getValue(s))
	}
	return maximal
}
