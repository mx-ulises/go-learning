func getSubsetXorSum(nums []int, i, current int) int {
	if i == len(nums) {
		return current
	}
	num := nums[i]
	i++
	total := getSubsetXorSum(nums, i, current)
	total += getSubsetXorSum(nums, i, current^num)
	return total
}

func subsetXORSum(nums []int) int {
	return getSubsetXorSum(nums, 0, 0)
}
