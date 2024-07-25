func GetDistance(x1, y1, x2, y2 int) int {
	x := x1 - x2
	y := y1 - y2
	return x*x + y*y
}

func CountPointsInside(query *[]int, points *[][]int) int {
	count := 0
	centerX, centerY, radio := (*query)[0], (*query)[1], (*query)[2]
	radio *= radio
	for _, point := range *points {
		distance := GetDistance(centerX, centerY, point[0], point[1])
		if distance <= radio {
			count++
		}
	}
	return count
}

func countPoints(points [][]int, queries [][]int) []int {
	output := []int{}
	for _, query := range queries {
		output = append(output, CountPointsInside(&query, &points))
	}
	return output
}
