package leetcode

func getArea(a, b, c []int) float64 {
	A := a[0] * (b[1] - c[1])
	B := b[0] * (c[1] - a[1])
	C := c[0] * (a[1] - b[1])
	factor := float64(A + B + C)
	if factor < 0 {
		factor = -factor
	}
	return factor / 2
}

func largestTriangleArea(points [][]int) float64 {
	area := 0.0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				candidate := getArea(points[i], points[j], points[k])
				area = max(area, candidate)
			}
		}
	}
	return area
}
