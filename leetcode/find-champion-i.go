func sum(team []int) int {
	total := 0
	for _, x := range team {
		total += x
	}
	return total
}

func findChampion(grid [][]int) int {
	maximal := len(grid) - 1
	for i, team := range grid {
		if sum(team) == maximal {
			return i
		}
	}
	return -1
}
