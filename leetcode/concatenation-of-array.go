func getConcatenation(nums []int) []int {
	output := []int{}
	output = append(output, nums...)
	output = append(output, nums...)
	return output
}
