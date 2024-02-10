func groupThePeople(groupSizes []int) [][]int {
	groupSize := map[int][]int{}
	groups := [][]int{}
	for i, size := range groupSizes {
		groupSize[size] = append(groupSize[size], i)
		if size == len(groupSize[size]) {
			groups = append(groups, groupSize[size])
			groupSize[size] = []int{}
		}
	}
	return groups
}
