func singleNumber(nums []int) []int {
	xor := 0
	for _, x := range nums {
		xor ^= x
	}
	diff := xor & -xor // Find a single bit of difference
	output := []int{0, 0}
	for _, x := range nums {
		if diff&x == 0 {
			output[0] ^= x
		} else {
			output[1] ^= x
		}
	}
	return output
}
