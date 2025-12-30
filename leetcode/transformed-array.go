package leetcode

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = getResultValue(i, nums)
	}
	return result
}

func getResultValue(i int, nums []int) int {
	j, n := (i + nums[i]), len(nums)
	for j < 0 {
		j += n
	}
	for n <= j {
		j -= n
	}
	return nums[j]
}
