package leetcode

func minOperations(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			return 1
		}
	}
	return 0
}
