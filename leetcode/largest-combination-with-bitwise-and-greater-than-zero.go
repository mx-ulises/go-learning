func updateMap(bitmap *map[int]int, candidate int) {
	bit := 1
	for bit <= candidate {
		if (candidate & bit) != 0 {
			(*bitmap)[bit]++
		}
		bit <<= 1
	}
}

func largestCombination(candidates []int) int {
	bitmap := map[int]int{}
	for _, candidate := range candidates {
		updateMap(&bitmap, candidate)
	}
	maximal := 0
	for _, count := range bitmap {
		maximal = max(maximal, count)
	}
	return maximal
}
