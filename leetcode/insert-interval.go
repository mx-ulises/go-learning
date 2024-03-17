func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func insert(intervals [][]int, newInterval []int) [][]int {
	newIntervals := [][]int{}
	merged := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] || newInterval[1] < interval[0] {
			if newInterval[1] < interval[0] && merged == false {
				newIntervals = append(newIntervals, newInterval)
				merged = true
			}
			newIntervals = append(newIntervals, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	if merged == false {
		newIntervals = append(newIntervals, newInterval)
	}
	return newIntervals
}
