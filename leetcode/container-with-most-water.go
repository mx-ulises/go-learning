func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxAreaStack(height []int) int {
	s := [][]int{{0, height[0]}}
	maximalArea := 0
	for i := 1; i < len(height); i++ {
		for j := 0; j < len(s); j++ {
			x := i - s[j][0]
			y := min(height[i], s[j][1])
			maximalArea = max(maximalArea, x*y)
		}
		if s[len(s)-1][1] < height[i] {
			s = append(s, []int{i, height[i]})
		}
	}
	return maximalArea
}

func maxAreaPointers(height []int) int {
	start := 0
	end := len(height) - 1
	maximalArea := 0
	for start < end {
		n := end - start
		m := min(height[start], height[end])
		maximalArea = max(maximalArea, n*m)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return maximalArea
}

func maxArea(height []int) int {
	return maxAreaPointers(height)
}
