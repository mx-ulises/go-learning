func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosest(x int, y int, z int) int {
	p1 := abs(x - z)
	p2 := abs(y - z)
	switch {
	case p1 < p2:
		return 1
	case p2 < p1:
		return 2
	}
	return 0
}
