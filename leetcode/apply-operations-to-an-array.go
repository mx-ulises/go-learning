func findNextZero(i int, nums []int) int {
	for i < len(nums) && nums[i] != 0 {
		i++
	}
	return i
}

func findNextNonZero(j int, nums []int) int {
	for j < len(nums) && nums[j] == 0 {
		j++
	}
	return j
}

func applyOperations(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			nums[i-1] *= 2
			nums[i] = 0
		}
	}
	i := findNextZero(0, nums)
	j := i
	for j < n {
		j = findNextNonZero(j, nums)
		if j < n {
			nums[i] = nums[j]
			nums[j] = 0
		}
		i = findNextZero(i, nums)
	}
	return nums
}
