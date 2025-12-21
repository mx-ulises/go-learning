package leetcode

func mirrorDistance(n int) int {
	reversed := reverse(n)
	return abs(n - reversed)
}

func reverse(n int) int {
	reversed := 0
	for 0 < n {
		reversed = (10 * reversed) + (n % 10)
		n /= 10
	}
	return reversed
}

func abs(x int) int {
	return max(x, -x)
}
