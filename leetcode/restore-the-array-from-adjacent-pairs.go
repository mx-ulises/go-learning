func restoreArray(adjacentPairs [][]int) []int {
	neighbors := map[int][]int{}
	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		neighbors[u] = append(neighbors[u], v)
		neighbors[v] = append(neighbors[v], u)
	}
	s, n := []int{}, 2
	for u, _ := range neighbors {
		if len(neighbors[u]) == 1 {
			s = append(s, u)
		}
	}
	output, i := make([]int, len(neighbors)), 0
	visited := map[int]bool{}
	for 0 < n {
		n--
		u := s[n]
		s = s[:n]
		if visited[u] {
			continue
		}
		visited[u] = true
		output[i] = u
		i++
		for _, v := range neighbors[u] {
			s = append(s, v)
			n++
		}
	}
	return output
}
