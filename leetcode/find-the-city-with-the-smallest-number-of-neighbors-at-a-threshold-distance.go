func BuildGraph(n int, edges [][]int) *map[int][][]int {
	graph := map[int][][]int{}
	for _, edge := range edges {
		start, end, weight := edge[0], edge[1], edge[2]
		graph[start] = append(graph[start], []int{end, weight})
		graph[end] = append(graph[end], []int{start, weight})
	}
	return &graph
}

func GetNeighbors(city int, graph *map[int][][]int, distanceThreshold int) int {
	visited := map[int]int{}
	q := [][]int{[]int{city, 0}}
	for 0 < len(q) {
		start, distance := q[0][0], q[0][1]
		q = q[1:]
		_, ok := visited[start]
		if (distanceThreshold < distance) || (ok && visited[start] < distance) {
			continue
		}
		visited[start] = distance
		for i := 0; i < len((*graph)[start]); i++ {
			end, weigth := (*graph)[start][i][0], (*graph)[start][i][1]
			newDistance := distance + weigth
			q = append(q, []int{end, newDistance})
		}
	}
	return len(visited) - 1
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	graph := BuildGraph(n, edges)
	bestCity, minNeighbors := -1, n
	for city := 0; city < n; city++ {
		neighbors := GetNeighbors(city, graph, distanceThreshold)
		if neighbors <= minNeighbors {
			minNeighbors = neighbors
			bestCity = city
		}
	}
	return bestCity
}
