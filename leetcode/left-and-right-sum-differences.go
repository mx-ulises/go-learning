func leftRightDifference(nums []int) []int {
	right, left := 0, 0
	for _, num := range nums {
		right += num
	}
	output := make([]int, len(nums))
	for i, num := range nums {
		right -= num
		output[i] = left - right
		if output[i] < 0 {
			output[i] = right - left
		}
		left += num
	}
	return output
}
