func findMinArrowShots(points [][]int) int {
	sorter := func(i, j int) bool {
		return points[i][0] < points[j][0]
	}
	sort.Slice(points, sorter)
	end := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		if end < points[i][0] {
			count++
			end = points[i][1]
		} else if points[i][1] < end {
			end = points[i][1]
		}
	}
	return count
}
