func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	memory := make([]int, n+1)
	for i := 0; i < n; i++ {
		currentWidth, currentHeight := 0, 0
		j := i
		memory[i+1] = 10000001
		for 0 <= j {
			w, h := books[j][0], books[j][1]
			currentHeight = max(currentHeight, h)
			currentWidth += w
			if shelfWidth < currentWidth {
				break
			}
			memory[i+1] = min(memory[i+1], currentHeight+memory[j])
			j--
		}
	}
	return memory[n]
}
