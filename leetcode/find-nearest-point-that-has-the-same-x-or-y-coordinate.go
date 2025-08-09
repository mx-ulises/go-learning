package leetcode

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func isValid(x1 int, y1 int, x2 int, y2 int) bool {
	return x1 == x2 || y1 == y2
}

func nearestValidPoint(x int, y int, points [][]int) int {
	output := -1
	minDistance := 100000
	for i, point := range points {
		a, b := point[0], point[1]
		distance := manhattanDistance(x, y, a, b)
		if isValid(x, y, a, b) && distance < minDistance {
			minDistance = distance
			output = i
		}
	}
	return output
}
