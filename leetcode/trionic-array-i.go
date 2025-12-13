package leetcode

func isTrionic(nums []int) bool {
	flips := 0
	prev := 0
	if nums[1] <= nums[0] {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if nums[prev] == nums[i] {
			return false
		}
		if flipRequired(flips, nums[i], nums[prev]) {
			flips++
		}
		prev = i
	}
	return flips == 3
}

func flipRequired(flips int, current int, previous int) bool {
	if isDecreasing(flips) && previous < current {
		return true
	}
	return isIncreasing(flips) && current < previous
}

func isDecreasing(flips int) bool {
	return flips&1 == 0
}

func isIncreasing(flips int) bool {
	return !isDecreasing(flips)
}
