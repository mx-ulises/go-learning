func twoSum(nums []int, target int) []int {
	previousNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		b := target - a
		j, ok := previousNums[b]
		if ok {
			return []int{j, i}
		}
		previousNums[a] = i
	}
	return []int{}
}
