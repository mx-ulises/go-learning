package leetcode

func areNumbersAscending(s string) bool {
	prev, current := 0, 0
	for i := 0; i < len(s); i++ {
		c := int(s[i] - '0')
		if 0 <= c && c <= 9 {
			current = current*10 + c
		} else if current != 0 {
			if current <= prev {
				return false
			}
			prev, current = current, 0
		}
	}
	return current == 0 || prev < current
}
