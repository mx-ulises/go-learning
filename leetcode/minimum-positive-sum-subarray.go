package leetcode

func getMinimumSumSubarray(nums []int, size int) int {
	if len(nums) < size {
		return -1
	}
	output := 0
	for i := 0; i < size; i++ {
		output += nums[i]
	}
	current := output
	for i := size; i < len(nums); i++ {
		current = current + nums[i] - nums[i-size]
		if 0 < current {
			if output <= 0 {
				output = current
			} else {
				output = min(output, current)
			}
		}
	}
	return output
}

func minimumSumSubarray(nums []int, l int, r int) int {
	output := -1
	for i := l; i <= r; i++ {
		candidate := getMinimumSumSubarray(nums, i)
		if 0 < candidate {
			if output == -1 {
				output = candidate
			} else {
				output = min(output, candidate)
			}
		}
	}
	return output
}
