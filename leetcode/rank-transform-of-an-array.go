func GetSortedCopy(arr *[]int) *[]int {
	sorted := make([]int, len(*arr))
	copy(sorted, *arr)
	sort.Ints(sorted)
	return &sorted
}

func GetRankMap(sorted *[]int) *map[int]int {
	rankMap := map[int]int{}
	currentRank := 1
	for _, num := range *sorted {
		_, ok := rankMap[num]
		if !ok {
			rankMap[num] = currentRank
			currentRank++
		}
	}
	return &rankMap
}

func arrayRankTransform(arr []int) []int {
	sorted := GetSortedCopy(&arr)
	rankMap := GetRankMap(sorted)
	output := make([]int, len(arr))
	for i, num := range arr {
		output[i] = (*rankMap)[num]
	}
	return output
}
