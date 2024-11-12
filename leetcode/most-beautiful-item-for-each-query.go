func maximumBeauty(items [][]int, queries []int) []int {
	// Add query values just in case they are not in items
	for _, query := range queries {
		items = append(items, []int{query, 0})
	}
	// Sort the items by price
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	// Build a map with all possible prices getting the maximal beauty
	beautyMap := map[int]int{}
	currentMaxBeauty := 0
	for _, item := range items {
		currentMaxBeauty = max(currentMaxBeauty, item[1])
		beautyMap[item[0]] = currentMaxBeauty
	}
	// Build output array using queries and beauty map
	output := make([]int, len(queries))
	for i, query := range queries {
		output[i] = beautyMap[query]
	}
	return output
}
