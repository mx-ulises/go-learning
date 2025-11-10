package leetcode

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiagonal, maxArea := 0, 0
	for i := 0; i < len(dimensions); i++ {
		length, width := dimensions[i][0], dimensions[i][1]
		diagonal, area := getDiagonal(length, width), getArea(length, width)
		if maxDiagonal < diagonal {
			maxDiagonal, maxArea = diagonal, area
		} else if maxDiagonal == diagonal {
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func getDiagonal(length, width int) int {
	return length*length + width*width
}

func getArea(length, width int) int {
	return length * width
}
