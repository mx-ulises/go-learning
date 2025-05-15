func findKOr(nums []int, k int) int {
	bits := [32]int{}
	for _, num := range nums {
		i := 0
		for 0 < num {
			bits[i] += num & 1
			i++
			num >>= 1
		}
	}
	output, i := 0, 1
	for _, count := range bits {
		if k <= count {
			output |= i
		}
		i <<= 1
	}
	return output
}
