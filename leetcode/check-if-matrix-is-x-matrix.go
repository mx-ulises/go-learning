package leetcode

func checkXMatrix(grid [][]int) bool {
	j, k := 0, len(grid)-1
	for _, row := range grid {
		for i, x := range row {
			if i == j || i == k {
				if x == 0 {
					return false
				}
			} else if x != 0 {
				return false
			}
		}
		j++
		k--
	}
	return true
}
