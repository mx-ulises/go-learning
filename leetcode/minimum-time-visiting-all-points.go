func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0
	n := len(points)
	for i := 1; i < n; i++ {
		diffX := abs(points[i-1][0] - points[i][0])
		diffY := abs(points[i-1][1] - points[i][1])
		minTime += max(diffX, diffY)
	}
	return minTime
}
