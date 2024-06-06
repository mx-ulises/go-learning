func isNStraightHand2(hand []int, groupSize int) bool {
	cardCount := map[int]int{}
	for _, x := range hand {
		cardCount[x]++
	}
	cardOrder := []int{}
	for x, _ := range cardCount {
		cardOrder = append(cardOrder, x)
	}
	sort.Ints(cardOrder)
	for _, x := range cardOrder {
		aux := cardCount[x]
		if aux == 0 {
			continue
		}
		for j := 0; j < groupSize; j++ {
			cardCount[x+j] -= aux
			if cardCount[x+j] < 0 {
				return false
			}
		}
	}
	return true
}

func isNStraightHand(hand []int, groupSize int) bool {
	queue := [][]int{}
	sort.Ints(hand)
	i := 0
	for i < len(hand) {
		current := hand[i]
		j := 0
		for i < len(hand) && current == hand[i] {
			i++
			if j < len(queue) {
				if queue[j][len(queue[j])-1] == current-1 {
					queue[j] = append(queue[j], current)
				} else {
					return false
				}
			} else {
				queue = append(queue, []int{current})
			}
			j++
		}
		for 0 < len(queue) && len(queue[0]) == groupSize {
			queue = queue[1:]
		}
	}
	return len(queue) == 0
}
