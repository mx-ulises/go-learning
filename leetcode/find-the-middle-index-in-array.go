package leetcode

func findMiddleIndex(nums []int) int {
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}
	leftSum := 0
	for i, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}
