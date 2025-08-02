package leetcode

import "sort"

func splitNum(num int) int {
	digits := []int{}
	for 0 < num {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	a, b, i := 0, 0, 0
	for i < len(digits) {
		a = a*10 + digits[i]
		i++
		if i < len(digits) {
			b = b*10 + digits[i]
		}
		i++
	}
	return a + b
}
