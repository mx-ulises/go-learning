package leetcode

func numLen(x int) int {
	l := 0
	if x <= 0 {
		l++
	}
	for x != 0 {
		l++
		x /= 10
	}
	return l
}

func findColumnWidth(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[j] = max(ans[j], numLen(grid[i][j]))
		}
	}
	return ans
}
