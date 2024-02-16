func GetNumCount(arr []int) map[int]int {
	numCount := map[int]int{}
	for _, num := range arr {
		numCount[num]++
	}
	return numCount
}

func GetCounts(numCount map[int]int) []int {
	counts := []int{}
	for _, count := range numCount {
		counts = append(counts, count)
	}
	return counts
}

func findLeastNumOfUniqueInts(arr []int, k int) int {
	numCount := GetNumCount(arr)
	counts := GetCounts(numCount)
	slices.Sort(counts)
	for 0 < k && 0 < len(counts) {
		count := counts[0]
		if count <= k {
			counts = counts[1:]
		}
		k -= count
	}
	return len(counts)
}
