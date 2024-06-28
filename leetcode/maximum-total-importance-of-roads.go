func maximumImportance(n int, roads [][]int) int64 {
	pathCount := make([]int64, n)
	for i := 0; i < len(roads); i++ {
		pathCount[roads[i][0]] += 1
		pathCount[roads[i][1]] += 1
	}
	sort.Slice(pathCount, func(i, j int) bool { return pathCount[i] < pathCount[j] })
	var output int64
	for i := 0; i < n; i++ {
		output += pathCount[i] * (int64(i) + 1)
	}
	return output
}
