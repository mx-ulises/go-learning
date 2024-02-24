func numJewelsInStones(jewels string, stones string) int {
	jewelMap := map[rune]bool{}
	for _, jewel := range jewels {
		jewelMap[jewel] = true
	}
	jewelCount := 0
	for _, stone := range stones {
		if jewelMap[stone] {
			jewelCount++
		}
	}
	return jewelCount
}
