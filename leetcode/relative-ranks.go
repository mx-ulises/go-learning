func GetRelativeRank(position int) string {
	output := ""
	switch position {
	case 0:
		output = "Gold Medal"
	case 1:
		output = "Silver Medal"
	case 2:
		output = "Bronze Medal"
	default:
		output = strconv.Itoa(position + 1)
	}
	return output
}

func GetSortedPositions(score *[]int) *[]int {
	sortedPositions := make([]int, len(*score))
	copy(sortedPositions, *score)
	sort.Ints(sortedPositions)
	return &sortedPositions
}

func GetPositionMap(sortedPositions *[]int) *map[int]int {
	n := len(*sortedPositions)
	positionMap := map[int]int{}
	for i := 0; i < n; i++ {
		j := n - i - 1
		positionMap[(*sortedPositions)[j]] = i
	}
	return &positionMap
}

func findRelativeRanks(score []int) []string {
	sortedPositions := GetSortedPositions(&score)
	positionMap := GetPositionMap(sortedPositions)
	n := len(score)
	output := make([]string, n)
	for i, s := range score {
		output[i] = GetRelativeRank((*positionMap)[s])
	}
	return output
}
