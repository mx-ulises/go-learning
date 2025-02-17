func getTileCount(tiles string) map[byte]int {
	tileCount := map[byte]int{}
	for i := 0; i < len(tiles); i++ {
		tileCount[tiles[i]]++
	}
	return tileCount
}

func getTiles(current []byte, visited map[string]bool, tileCount map[byte]int) {
	s := string(current)
	if visited[s] {
		return
	}
	visited[s] = true
	for tile, count := range tileCount {
		if count == 0 {
			continue
		}
		current = append(current, tile)
		tileCount[tile]--
		getTiles(current, visited, tileCount)
		tileCount[tile]++
		current = current[:len(current)-1]
	}
}

func numTilePossibilities(tiles string) int {
	tileCount := getTileCount(tiles)
	current := []byte{}
	visited := map[string]bool{}
	getTiles(current, visited, tileCount)
	return len(visited) - 1
}
