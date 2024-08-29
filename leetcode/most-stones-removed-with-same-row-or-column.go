func BuildGraphs(stones *[][]int) (*map[int][]int, *map[int][]int) {
	graphX := map[int][]int{}
	graphY := map[int][]int{}
	for i, stone := range *stones {
		x, y := stone[0], stone[1]
		graphX[x] = append(graphX[x], i)
		graphY[y] = append(graphY[y], i)
	}
	return &graphX, &graphY
}

func GetCount(stones *[][]int, graphX, graphY *map[int][]int, i int, visited *map[int]bool) int {
	if (*visited)[i] {
		return 0
	}
	(*visited)[i] = true
	count := 1
	x, y := (*stones)[i][0], (*stones)[i][1]
	for _, j := range (*graphX)[x] {
		count += GetCount(stones, graphX, graphY, j, visited)
	}
	for _, j := range (*graphY)[y] {
		count += GetCount(stones, graphX, graphY, j, visited)
	}
	return count
}

func removeStones(stones [][]int) int {
	visited := map[int]bool{}
	count := 0
	graphX, graphY := BuildGraphs(&stones)
	for i, _ := range stones {
		count += max(GetCount(&stones, graphX, graphY, i, &visited)-1, 0)
	}
	return count
}
