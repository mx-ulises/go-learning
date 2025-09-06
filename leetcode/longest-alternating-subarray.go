package leetcode

func alternatingSubarray(nums []int) int {
	expected := 1
	currentSize := 1
	maximalSize := 1
	diff := 0
	for i := 1; i < len(nums); i++ {
		diff = nums[i] - nums[i-1]
		if diff == expected {
			expected = -expected
			currentSize++
		} else if diff == 1 {
			expected = -1
			currentSize = 2
		} else {
			expected = 1
			currentSize = 1
		}
		maximalSize = max(maximalSize, currentSize)
	}
	if maximalSize == 1 {
		return -1
	}
	return maximalSize
}
