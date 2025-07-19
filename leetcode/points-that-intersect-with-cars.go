package leetcode

import "sort"

func numberOfPoints(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	points := 0
	start, end := nums[0][0], nums[0][1]
	for _, car := range nums {
		if car[0] <= end {
			end = max(end, car[1])
		} else {
			points += end - start + 1
			start, end = car[0], car[1]
		}
	}
	points += end - start + 1
	return points
}
