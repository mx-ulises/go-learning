func initGraph(n int) [][]int {
	graph := make([][]int, n+1)
	for i := 0; i < n; i++ {
		graph[i] = []int{i + 1}
	}
	return graph
}

func initMinDistances(n int) []int {
	n++
	minDistances := make([]int, n)
	for i := 0; i < n; i++ {
		minDistances[i] = i
	}
	return minDistances
}

func updateMinDistances(graph [][]int, minDistances []int, start, n int) {
	q := [][]int{{minDistances[start], start}}
	visited := map[int]bool{}
	for 0 < len(q) {
		d, s := q[0][0], q[0][1]
		q = q[1:]
		minDistances[s] = min(minDistances[s], d)
		if s == n {
			return
		}
		if visited[s] {
			continue
		}
		visited[s] = true
		d++
		for _, t := range graph[s] {
			q = append(q, []int{d, t})
		}
	}
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := initGraph(n)
	minDistances := initMinDistances(n)
	output := make([]int, len(queries))
	n -= 1
	for i, query := range queries {
		start, end := query[0], query[1]
		graph[start] = append(graph[start], end)
		minDistances[end] = min(minDistances[end], minDistances[start]+1)
		updateMinDistances(graph, minDistances, end, n)
		output[i] = minDistances[n]
	}
	return output
}
