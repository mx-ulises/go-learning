func createTargetArray(nums []int, index []int) []int {
	n := len(nums)
	output := []int{}
	for i := 0; i < n; i++ {
		if index[i] == len(output) {
			output = append(output, nums[i])
			continue
		}
		new_output := []int{}
		for j := 0; j < len(output); j++ {
			if j == index[i] {
				new_output = append(new_output, nums[i])
			}
			new_output = append(new_output, output[j])
		}
		output = new_output
	}
	return output
}
