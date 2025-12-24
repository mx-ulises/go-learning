package leetcode

func maxHeightOfTriangle(red int, blue int) int {
	heightWithRed := maxHeightOfTriangleStartingWith(red, blue)
	heightWithBlue := maxHeightOfTriangleStartingWith(blue, red)
	return max(heightWithRed, heightWithBlue)
}

func maxHeightOfTriangleStartingWith(first int, second int) int {
	h := 1
	for h <= first {
		first, second = second, first-h
		h++
	}
	return h - 1
}
