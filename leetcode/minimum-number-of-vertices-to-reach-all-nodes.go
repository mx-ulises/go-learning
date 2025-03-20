func findSmallestSetOfVertices(n int, edges [][]int) []int {
	touched := make([]bool, n)
	for _, edge := range edges {
		touched[edge[1]] = true
	}
	output := []int{}
	for i, count := range touched {
		if count == false {
			output = append(output, i)
		}
	}
	return output
}
