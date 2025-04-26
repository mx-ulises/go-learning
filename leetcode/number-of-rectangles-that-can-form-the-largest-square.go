func countGoodRectangles(rectangles [][]int) int {
	maximal, count := 0, 0
	for _, rectangle := range rectangles {
		l := min(rectangle[0], rectangle[1])
		if maximal < l {
			maximal, count = l, 0
		}
		if maximal == l {
			count++
		}
	}
	return count
}
