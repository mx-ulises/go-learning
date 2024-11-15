func fillTotalWeight(totalWeight map[int]int, items [][]int) {
	for _, item := range items {
		value, weight := item[0], item[1]
		totalWeight[value] += weight
	}
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	totalWeight := map[int]int{}
	fillTotalWeight(totalWeight, items1)
	fillTotalWeight(totalWeight, items2)
	output := make([][]int, len(totalWeight))
	i := 0
	for value, weight := range totalWeight {
		output[i] = append(output[i], value, weight)
		i++
	}
	sort.Slice(output, func(i, j int) bool {
		return output[i][0] < output[j][0]
	})
	return output
}
