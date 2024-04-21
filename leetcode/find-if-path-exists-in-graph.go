func GetGraph(edges *[][]int) *map[int][]int {
	graph := map[int][]int{}
	for _, edge := range *edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return &graph
}

func IsReachable(n int, graph *map[int][]int, source int, destination int) bool {
	visited := make([]bool, n)
	queue := []int{source}
	for 0 < len(queue) {
		node := queue[0]
		queue = queue[1:]
		if visited[node] {
			continue
		}
		if node == destination {
			return true
		}
		queue = append(queue, (*graph)[node]...)
		visited[node] = true
	}
	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := GetGraph(&edges)
	return IsReachable(n, graph, source, destination)
}
