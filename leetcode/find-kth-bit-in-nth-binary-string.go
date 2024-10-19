func Find(n int, k int) int {
	if n == 1 {
		return 0
	}
	mid := (1 << n) / 2
	if k < mid {
		return Find(n-1, k)
	}
	if mid < k {
		k -= 2 * mid
		return Find(n-1, -k) ^ 1
	}
	return 1
}

func findKthBit(n int, k int) byte {
	if Find(n, k) == 0 {
		return '0'
	}
	return '1'
}
