package leetcode

func getMax(nums []int) int {
	maximal := -1
	for _, num := range nums {
		maximal = max(maximal, num)
	}
	return maximal
}

func largestInteger(nums []int, k int) int {
	n := len(nums)
	if k == n {
		return getMax(nums)
	}
	candidateCount := map[int]int{}
	for _, num := range nums {
		if k == 1 || num == nums[0] || num == nums[n-1] {
			candidateCount[num]++
		}
	}
	output := -1
	for num, count := range candidateCount {
		if count == 1 {
			output = max(num, output)
		}
	}
	return output
}
