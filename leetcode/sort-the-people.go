func sortPeople(names []string, heights []int) []string {
	heightToNameIndex := map[int]int{}
	for i, h := range heights {
		heightToNameIndex[h] = i
	}
	sort.Ints(heights)
	n := len(names)
	orderedNames := make([]string, n)
	for i, h := range heights {
		orderedNames[n-i-1] = names[heightToNameIndex[h]]
	}
	return orderedNames
}
