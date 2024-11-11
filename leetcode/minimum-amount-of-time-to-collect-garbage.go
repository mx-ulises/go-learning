func updateGarbage(house string, g, p, m, new int) (int, int, int) {
	for _, c := range house {
		switch c {
		case 'G':
			g = new
		case 'P':
			p = new
		case 'M':
			m = new
		}
	}
	return g, p, m
}

func sum(travels []int) int {
	totalTime := 0
	for _, t := range travels {
		totalTime += t
	}
	return totalTime
}

func garbageCollection(garbage []string, travel []int) int {
	g, p, m := 0, 0, 0
	totalTime := 0
	for i := 0; i < len(garbage); i++ {
		totalTime += len(garbage[i])
		g, p, m = updateGarbage(garbage[i], g, p, m, i)
	}
	totalTime += sum(travel[:g]) + sum(travel[:p]) + sum(travel[:m])
	return totalTime
}
