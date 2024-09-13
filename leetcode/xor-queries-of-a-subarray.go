func xorQueries(arr []int, queries [][]int) []int {
	n, m := len(arr), len(queries)
	xors := make([]int, n+1)
	for i := 0; i < n; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}
	output := make([]int, m)
	for i := 0; i < m; i++ {
		l, r := queries[i][0], queries[i][1]+1
		output[i] = xors[l] ^ xors[r]
	}
	return output
}
