func getGroup(x int) int {
	group := 0
	for 0 < x {
		group += x % 10
		x /= 10
	}
	return group
}

func countLargestGroup(n int) int {
	maxGroupSize := 0
	maxGroupSizeCount := 0
	groupSizes := [37]int{}
	for x := 1; x <= n; x++ {
		group := getGroup(x)
		groupSizes[group]++
		if maxGroupSize < groupSizes[group] {
			maxGroupSize = groupSizes[group]
			maxGroupSizeCount = 0
		}
		if maxGroupSize == groupSizes[group] {
			maxGroupSizeCount++
		}
	}
	return maxGroupSizeCount
}
