func queryResults(limit int, queries [][]int) []int {
	limit++
	colorCount := 0
	colorBalls := map[int]int{0: limit}
	ballColors := map[int]int{}
	output := make([]int, len(queries))
	for i, query := range queries {
		ball, newColor := query[0], query[1]
		// Remove previous color
		currentColor := ballColors[ball]
		colorBalls[currentColor]--
		if currentColor != 0 && colorBalls[currentColor] == 0 {
			colorCount--
		}
		// Add new color
		colorBalls[newColor]++
		if colorBalls[newColor] == 1 {
			colorCount++
		}
		// Update color of the ball and add output
		ballColors[ball] = newColor
		output[i] = colorCount
	}
	return output
}
