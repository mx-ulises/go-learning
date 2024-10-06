func GetPaths(graph *[][]int, paths *[][]int, start int, end int, current *[]int) {
	if start == end {
		path := make([]int, len(*current))
		copy(path, *current)
		*paths = append(*paths, path)
	} else {
		for _, succ := range (*graph)[start] {
			*current = append(*current, succ)
			GetPaths(graph, paths, succ, end, current)
			*current = (*current)[:len(*current)-1]
		}
	}

}

func allPathsSourceTarget(graph [][]int) [][]int {
	output := [][]int{}
	GetPaths(&graph, &output, 0, len(graph)-1, &[]int{0})
	return output
}
