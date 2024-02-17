func numIdenticalPairs(nums []int) int {
	goodPairs := 0
	numCount := map[int]int{}
	for _, num := range nums {
		goodPairs += numCount[num]
		numCount[num]++
	}
	return goodPairs
}
