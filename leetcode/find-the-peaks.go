func findPeaks(mountain []int) []int {
	output := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i-1] < mountain[i] && mountain[i+1] < mountain[i] {
			output = append(output, i)
		}
	}
	return output
}
