func maxCoins(piles []int) int {
	n, i, count := len(piles)/3, len(piles)-2, 0
	sort.Ints(piles)
	for 0 < n {
		count += piles[i]
		i -= 2
		n--
	}
	return count
}
