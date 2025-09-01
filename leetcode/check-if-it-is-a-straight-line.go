package leetcode

import "sort"

func checkStraightLine(coordinates [][]int) bool {
	// Sort based on the first element of each sub-array
	sort.Slice(coordinates, func(i, j int) bool {
		return coordinates[i][0] < coordinates[j][0]
	})
	m := float64(coordinates[0][0]-coordinates[1][0]) / float64(coordinates[0][1]-coordinates[1][1])
	n := len(coordinates)
	for i := 1; i < n; i++ {
		j := i - 1
		candidate := float64(coordinates[j][0]-coordinates[i][0]) / float64(coordinates[j][1]-coordinates[i][1])
		if candidate != m {
			return false
		}
	}
	return true
}
