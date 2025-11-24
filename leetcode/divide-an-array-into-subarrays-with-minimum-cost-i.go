package leetcode

func minimumCost(nums []int) int {
	first := nums[0]
	second, third := getSecondAndThirdSubarrays(nums)
	return first + second + third
}

func getSecondAndThirdSubarrays(nums []int) (int, int) {
	second, third := 1000, 1000
	for i := 1; i < len(nums); i++ {
		second, third = checkAndUpdate(second, third, nums[i])
	}
	return second, third
}

func checkAndUpdate(second int, third int, candidate int) (int, int) {
	if candidate < second {
		third, second = second, candidate
	} else if candidate < third {
		third = candidate
	}
	return second, third
}
