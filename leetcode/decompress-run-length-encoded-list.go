func decompressRLElist(nums []int) []int {
	output := []int{}
	n := len(nums) / 2
	for i := 0; i < n; i++ {
		j := i * 2
		freq, num := nums[j], nums[j+1]
		for 0 < freq {
			output = append(output, num)
			freq--
		}
	}
	return output
}
