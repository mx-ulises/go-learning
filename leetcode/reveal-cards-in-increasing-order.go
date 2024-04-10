func GetOrder(n int) []int {
	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i
	}
	newOrder := make([]int, 0)
	for 0 < len(order) {
		newOrder = append(newOrder, order[0])
		order = order[1:]
		if 0 < len(order) {
			order = append(order, order[0])
			order = order[1:]
		}
	}
	return newOrder
}

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	sort.Ints(deck)
	order := GetOrder(n)
	output := make([]int, n)
	for i, j := range order {
		output[j] = deck[i]
	}
	return output
}
