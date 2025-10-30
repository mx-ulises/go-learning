package leetcode

func removeZeros(n int64) int64 {
	var output, m, d int64 = 0, 1, -1
	for 0 < n {
		d, n = n%10, n/10
		if 0 < d {
			output += m * d
			m *= 10
		}
	}
	return output
}
