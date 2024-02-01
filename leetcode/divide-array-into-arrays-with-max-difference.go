func divideArray(nums []int, k int) [][]int {
	slices.Sort(nums)
	output := [][]int{}
	output = append(output, []int{nums[0]})
	j := 0
	for i := 1; i < len(nums); i++ {
		if len(output[j]) == 3 {
			output = append(output, []int{nums[i]})
			j++
		} else {
			if k < (nums[i] - output[j][0]) {
				return make([][]int, 0)
			}
			output[j] = append(output[j], nums[i])
		}
	}
	return output
}
