func rearrangeArray(nums []int) []int {
	p, n := 0, 1
	output := make([]int, len(nums))
	for _, num := range nums {
		if 0 < num {
			output[p] = num
			p += 2
		} else {
			output[n] = num
			n += 2
		}
	}
	return output
}
