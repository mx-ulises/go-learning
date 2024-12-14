func getRods(rings string) map[byte]int {
	rods := map[byte]int{}
	ringMask := map[byte]int{'R': 0b001, 'G': 0b010, 'B': 0b100}
	i, n := 0, len(rings)
	for i < n {
		color := rings[i]
		rod := rings[i+1]
		rods[rod] |= ringMask[color]
		i += 2
	}
	return rods
}

func countPoints(rings string) int {
	rods := getRods(rings)
	fullRods := 0
	for _, rodColors := range rods {
		if rodColors == 0b111 {
			fullRods++
		}
	}
	return fullRods
}
