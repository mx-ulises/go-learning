package leetcode

func countSubarrays(nums []int) int {
	subarrayCount := 0
	first, second := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		third := nums[i]
		if isValid(first, second, third) {
			subarrayCount++
		}
		first, second = second, third
	}
	return subarrayCount
}

func isValid(first int, second int, third int) bool {
	return 2*(first+third) == second
}
