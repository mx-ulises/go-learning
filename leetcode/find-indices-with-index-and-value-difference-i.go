package leetcode

func abs(x int) int {
	return max(x, -x)
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	for i := 0; i < len(nums)-indexDifference; i++ {
		for j := i + indexDifference; j < len(nums); j++ {
			if valueDifference <= abs(nums[i]-nums[j]) {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
