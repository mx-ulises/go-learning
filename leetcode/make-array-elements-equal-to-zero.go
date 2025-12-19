package leetcode

func countValidSelections(nums []int) int {
	leftSum, rightSum := 0, getSum(nums)
	validSelections := 0
	for i := 0; i < len(nums); i++ {
		validSelections += countValidSelectionsInPosition(nums[i], leftSum, rightSum)
		leftSum, rightSum = updateSums(nums[i], leftSum, rightSum)
	}
	return validSelections
}

func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func countValidSelectionsInPosition(num int, leftSum int, rightSum int) int {
	if num != 0 {
		return 0
	}
	diff := abs(leftSum - rightSum)
	return max(2-diff, 0)
}

func abs(num int) int {
	return max(num, -num)
}

func updateSums(num int, leftSum int, rightSum int) (int, int) {
	return leftSum + num, rightSum - num
}
