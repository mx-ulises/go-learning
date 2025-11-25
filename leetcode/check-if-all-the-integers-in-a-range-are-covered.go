package leetcode

import "sort"

func isCovered(ranges [][]int, left int, right int) bool {
	sortRangesByStart(ranges)
	for _, r := range ranges {
		left = checkIfValueInRangeAndMove(r, left)
		if allValuesCovered(left, right) {
			return true
		}

	}
	return false
}

func sortRangesByStart(ranges [][]int) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
}

func checkIfValueInRangeAndMove(r []int, value int) int {
	if checkIfValueInRange(r[0], r[1], value) {
		value = r[1] + 1
	}
	return value
}

func checkIfValueInRange(start int, end int, value int) bool {
	return start <= value && value <= end
}

func allValuesCovered(left int, right int) bool {
	return right < left
}
