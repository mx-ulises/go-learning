func shuffle(nums []int, n int) []int {
	output := make([]int, 2*n)
	for i := 0; i < n; i++ {
		j := i * 2
		output[j] = nums[i]
		output[j+1] = nums[i+n]
	}
	return output
}
