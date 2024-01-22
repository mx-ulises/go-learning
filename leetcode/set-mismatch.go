func findErrorNums(nums []int) []int {
	output := make([]int, 2)
	numSet := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		numSet[nums[i]]++
	}
	for i := 1; i < len(numSet); i++ {
		if numSet[i] == 0 {
			output[1] = i
		}
		if numSet[i] == 2 {
			output[0] = i
		}
	}
	return output
}
