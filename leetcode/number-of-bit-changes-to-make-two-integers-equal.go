package leetcode

func minChanges(n int, k int) int {
	changes := 0
	for 0 < n {
		a, b := n&1, k&1
		if a == 1 && b == 0 {
			changes++
		}
		if a == 0 && b == 1 {
			return -1
		}
		n, k = n>>1, k>>1
	}
	if k == 0 {
		return changes
	}
	return -1
}
