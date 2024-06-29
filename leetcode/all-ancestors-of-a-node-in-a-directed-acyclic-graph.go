func BuildGraph(graph *[][]int, edges *[][]int) {
	for _, edge := range *edges {
		origin, destiny := edge[0], edge[1]
		(*graph)[origin] = append((*graph)[origin], destiny)
	}
}

func AddAncestor(graph *[][]int, origin int, ancestor int, ancestors *[][]int) {
	for _, destiny := range (*graph)[origin] {
		n := len((*ancestors)[destiny])
		if n == 0 || (*ancestors)[destiny][n-1] != ancestor {
			(*ancestors)[destiny] = append((*ancestors)[destiny], ancestor)
			AddAncestor(graph, destiny, ancestor, ancestors)
		}
	}
}

func getAncestors(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	BuildGraph(&graph, &edges)
	ancestors := make([][]int, n)
	for i := 0; i < n; i++ {
		AddAncestor(&graph, i, i, &ancestors)
	}
	return ancestors
}
