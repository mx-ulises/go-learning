package leetcode

func numberOfCuts(n int) int {
	if 1 < n && n&1 == 1 {
		return n
	}
	return n >> 1
}
