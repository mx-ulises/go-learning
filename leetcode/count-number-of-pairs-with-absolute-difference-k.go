func countKDifference(nums []int, k int) int {
	count := 0
	countMap := map[int]int{}
	for _, x := range nums {
		count += countMap[x+k] + countMap[x-k]
		countMap[x]++
	}
	return count
}
