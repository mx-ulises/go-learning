func maxChunksToSorted(arr []int) int {
	end := arr[0]
	chunks := 1
	for i := 1; i < len(arr); i++ {
		candidateEnd := max(i, arr[i])
		if min(i, arr[i]) <= end {
			end = max(end, candidateEnd)
		} else {
			end = candidateEnd
			chunks++
		}
	}
	return chunks
}
