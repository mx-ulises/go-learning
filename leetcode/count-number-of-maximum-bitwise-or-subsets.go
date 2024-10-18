func CountSubsets(nums *[]int, i int, current int, maximum int) int {
	if i == len(*nums) {
		if current == maximum {
			return 1
		}
		return 0
	}
	subsets := CountSubsets(nums, i+1, current, maximum)
	subsets += CountSubsets(nums, i+1, current|(*nums)[i], maximum)
	return subsets
}

func countMaxOrSubsets(nums []int) int {
	maximum := 0
	for _, num := range nums {
		maximum |= num
	}
	return CountSubsets(&nums, 0, 0, maximum)
}
