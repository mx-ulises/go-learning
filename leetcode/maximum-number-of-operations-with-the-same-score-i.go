package leetcode

func maxOperations(nums []int) int {
	score := nums[0] + nums[1]
	operationCount := 1
	for i := 2; i < len(nums)-1; i += 2 {
		candidate := nums[i] + nums[i+1]
		if candidate != score {
			return operationCount
		}
		operationCount++
	}
	return operationCount
}
