func maxWidthOfVerticalArea(points [][]int) int {
	n := len(points)
	x_axis := make([]int, n)
	for i := 0; i < n; i++ {
		x_axis[i] = points[i][0]
	}
	sort.Ints(x_axis)
	area := 0
	for i := 1; i < n; i++ {
		area = max(area, x_axis[i]-x_axis[i-1])
	}
	return area
}
