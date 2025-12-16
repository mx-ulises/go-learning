package leetcode

func minimumPairRemoval(nums []int) int {
	operations := 0
	for !isNonDecreasing(nums) {
		nums = removeMinimumPair(nums)
		operations++
	}
	return operations
}

func isNonDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}

func removeMinimumPair(nums []int) []int {
	minPairPosition := getMinPairPosition(nums)
	return shiftLeft(nums, minPairPosition)
}

func getMinPairPosition(nums []int) int {
	minPairPosition, minPairSum := -1, 50001
	for i := 0; i < len(nums)-1; i++ {
		minPairPosition, minPairSum = checkAndUpdateMinPair(minPairPosition, minPairSum, i, nums[i]+nums[i+1])
	}
	//fmt.Println("MinPairSum: ", minPairSum, ", MinPairPos: ", minPairPosition)
	return minPairPosition
}

func checkAndUpdateMinPair(minPairPosition int, minPairSum int, candidatePairPosition int, candidatePairSum int) (int, int) {
	if candidatePairSum < minPairSum {
		return candidatePairPosition, candidatePairSum
	}
	return minPairPosition, minPairSum
}

func shiftLeft(nums []int, position int) []int {
	for i := position; i < len(nums)-1; i++ {
		nums[i] += nums[i+1]
		nums[i+1] = 0
	}
	return nums[:len(nums)-1]
}
