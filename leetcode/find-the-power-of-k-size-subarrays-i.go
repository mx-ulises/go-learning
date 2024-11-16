func resultsArray(nums []int, k int) []int {
	m := len(nums) - k + 1
	output := make([]int, m)
	for i := 0; i < m; i++ {
		maximum := nums[i] - 1
		for j := 0; j < k; j++ {
			maximum++
			if maximum != nums[i+j] {
				maximum = -1
				break
			}
		}
		output[i] = maximum
	}
	return output
}
