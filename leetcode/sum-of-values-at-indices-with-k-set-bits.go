func GetBitCount(i int) int {
	bits := 0
	for 0 < i {
		bits += i & 1
		i >>= 1
	}
	return bits
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	total := 0
	for i, num := range nums {
		if GetBitCount(i) == k {
			total += num
		}
	}
	return total
}
