func frequencySort(s string) string {
	charClusters := map[rune][]rune{}
	for _, c := range s {
		charClusters[c] = append(charClusters[c], c)
	}
	clusters := [][]rune{}
	for _, cluster := range charClusters {
		clusters = append(clusters, cluster)
	}
	sort.Slice(clusters, func(i, j int) bool {
		return len(clusters[i]) < len(clusters[j])
	})
	output := []rune{}
	for i := len(clusters) - 1; 0 <= i; i-- {
		output = append(output, clusters[i]...)
	}
	return string(output)
}
