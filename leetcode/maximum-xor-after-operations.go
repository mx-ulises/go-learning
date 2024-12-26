func maximumXOR(nums []int) int {
	output := 0
	for _, num := range nums {
		output |= num
	}
	return output
}
