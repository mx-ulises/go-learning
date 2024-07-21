func GetOrder(k int, conditions [][]int) []int {
	startNumbers := map[int]bool{}
	for i := 1; i <= k; i++ {
		startNumbers[i] = true
	}
	pred := map[int][]int{}
	succ := map[int][]int{}
	for _, condition := range conditions {
		origin, destiny := condition[0], condition[1]
		startNumbers[destiny] = false
		succ[origin] = append(succ[origin], destiny)
		pred[destiny] = append(pred[destiny], origin)
	}
	q := []int{}
	for number, isStart := range startNumbers {
		if isStart {
			q = append(q, number)
		}
	}
	order := []int{}
	visited := map[int]bool{}
	for 0 < len(q) {
		number := q[0]
		q = q[1:]
		if visited[number] {
			continue
		}
		allPredVisited := true
		for _, prev := range pred[number] {
			if visited[prev] == false {
				allPredVisited = false
				break
			}
		}
		if allPredVisited == false {
			continue
		}
		visited[number] = true
		order = append(order, number)
		for _, suc := range succ[number] {
			q = append(q, suc)
		}
	}
	return order
}

func GetPositions(k int, rowOrder []int, colOrder []int) map[int][]int {
	positions := map[int][]int{}
	for number := 1; number <= k; number++ {
		positions[number] = []int{0, 0}
	}
	for i := 0; i < k; i++ {
		positions[rowOrder[i]][0] = i
		positions[colOrder[i]][1] = i
	}
	return positions
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := GetOrder(k, rowConditions)
	colOrder := GetOrder(k, colConditions)
	if len(rowOrder) < k || len(colOrder) < k {
		return [][]int{}
	}
	positions := GetPositions(k, rowOrder, colOrder)
	matrix := [][]int{}
	for i := 0; i < k; i++ {
		matrix = append(matrix, make([]int, k))
	}
	for number, position := range positions {
		matrix[position[0]][position[1]] = number
	}
	return matrix
}
