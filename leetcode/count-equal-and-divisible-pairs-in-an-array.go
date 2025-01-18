func countPairs(nums []int, k int) int {
	pairCount := 0
	numIndexes := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for _, j := range numIndexes[num] {
			if (i*j)%k == 0 {
				pairCount++
			}
		}
		numIndexes[num] = append(numIndexes[num], i)
	}
	return pairCount
}
