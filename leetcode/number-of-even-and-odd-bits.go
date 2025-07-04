package leetcode

func evenOddBit(n int) []int {
	output, i := make([]int, 2), 0
	for 0 < n {
		output[i] += n & 1
		n >>= 1
		i ^= 1
	}
	return output
}
