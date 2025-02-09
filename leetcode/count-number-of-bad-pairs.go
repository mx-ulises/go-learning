func countBadPairs(nums []int) int64 {
	var badPairs int64
	memory := map[int]int64{}
	for i, num := range nums {
		candidate := num - i
		badPairs += int64(i) - memory[candidate]
		memory[candidate]++
	}
	return badPairs
}
