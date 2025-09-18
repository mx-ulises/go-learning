package leetcode

import "sort"

func getParityToDigits(num int) ([2][]int, []int) {
	parityToDigits := [2][]int{}
	parityOrder := []int{}
	d, p := -1, -1
	for 0 < num {
		d, num = num%10, num/10
		p = d & 1
		parityToDigits[p] = append(parityToDigits[p], d)
		parityOrder = append(parityOrder, p)
	}
	sort.Ints(parityToDigits[0])
	sort.Ints(parityToDigits[1])
	return parityToDigits, parityOrder
}

func largestInteger(num int) int {
	parityToDigits, parityOrder := getParityToDigits(num)
	output := 0
	for i := len(parityOrder) - 1; 0 <= i; i-- {
		p := parityOrder[i]
		last := len(parityToDigits[p]) - 1
		d := parityToDigits[p][last]
		parityToDigits[p] = parityToDigits[p][:last]
		output = 10*output + d
	}
	return output
}
