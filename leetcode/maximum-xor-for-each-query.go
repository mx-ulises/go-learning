func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	output := make([]int, n)
	maximal := (1 << maximumBit) - 1
	current := 0
	for i := 0; i < n; i++ {
		current ^= nums[i]
		mask := (maximal & current) ^ maximal
		j := n - i - 1
		output[j] = mask
	}
	return output
}
