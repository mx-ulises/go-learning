const MAXIMAL = 1073741824 // 2^30

func isPowerOfTwo(n int) bool {
	if n <= 0 || MAXIMAL < n {
		return false
	} else if n == MAXIMAL {
		return true
	}
	return ((MAXIMAL - n) & n) == n
}
