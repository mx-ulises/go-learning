package leetcode

func minMoves(nums []int) int {
	maximal := getMaximal(nums)
	return getMovesToMakeEqualToMax(nums, maximal)
}

func getMaximal(nums []int) int {
	maximal := nums[0]
	for _, num := range nums {
		maximal = max(maximal, num)
	}
	return maximal
}

func getMovesToMakeEqualToMax(nums []int, maximal int) int {
	moves := 0
	for _, num := range nums {
		moves += maximal - num
	}
	return moves
}
