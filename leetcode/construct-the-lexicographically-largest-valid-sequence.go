func construct(output []int, visited map[int]bool, n, index int) bool {
	for index < len(output) && output[index] != 0 {
		index++
	}
	if index == len(output) {
		return true
	}
	for i := n; 0 < i; i-- {
		nextIndex := index + i
		if i == 1 {
			nextIndex = index
		}
		if visited[i] == false && nextIndex < len(output) && output[nextIndex] == 0 {
			visited[i] = true
			output[index] = i
			output[nextIndex] = i
			if construct(output, visited, n, index) {
				return true
			}
			output[nextIndex] = 0
			output[index] = 0
			visited[i] = false
		}
	}
	return false
}

func constructDistancedSequence(n int) []int {
	output := make([]int, n*2-1)
	visited := map[int]bool{}
	construct(output, visited, n, 0)
	return output
}
