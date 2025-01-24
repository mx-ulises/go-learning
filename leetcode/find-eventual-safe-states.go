func isSafe(graph [][]int, safe, unsafe, visited []bool, i int) bool {
	if safe[i] {
		return true
	}
	if unsafe[i] || visited[i] {
		return false
	}
	visited[i] = true
	for _, j := range graph[i] {
		if !isSafe(graph, safe, unsafe, visited, j) {
			unsafe[i] = true
			return false
		}
	}
	visited[i] = false
	safe[i] = true
	return true
}
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safe, unsafe, visited := make([]bool, n), make([]bool, n), make([]bool, n)
	output := []int{}
	for i := 0; i < n; i++ {
		if isSafe(graph, safe, unsafe, visited, i) {
			output = append(output, i)
		}
	}
	return output
}
