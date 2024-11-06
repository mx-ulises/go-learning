func getBits(num int) int {
	bits := 0
	for num != 0 {
		bits += num & 1
		num >>= 1
	}
	return bits
}

func canSortArray(nums []int) bool {
	prevMaximal := 0
	i, n := 0, len(nums)
	for i < n {
		bits := getBits(nums[i])
		maximal, minimal := nums[i], nums[i]
		for i < n && bits == getBits(nums[i]) {
			maximal = max(maximal, nums[i])
			minimal = min(minimal, nums[i])
			i++
		}
		if minimal < prevMaximal {
			return false
		}
		prevMaximal = maximal
	}
	return true
}
