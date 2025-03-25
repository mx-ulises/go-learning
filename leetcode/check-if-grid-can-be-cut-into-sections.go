func getPairs(rectangles [][]int, i, j int) [][]int {
	pairs := make([][]int, 0)
	for _, rectangle := range rectangles {
		start, end := rectangle[i], rectangle[j]
		pairs = append(pairs, []int{start, end})
	}
	return pairs
}

func sortPairs(pairs [][]int) {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})
}

func canCut(pairs [][]int, minCuts int) bool {
	cuts := 0
	end := pairs[0][1]
	for _, pair := range pairs {
		if pair[0] < end {
			end = max(end, pair[1])
		} else {
			cuts++
			end = pair[1]
		}
		if cuts == minCuts {
			return true
		}
	}
	return false
}

func checkCut(n int, rectangles [][]int, i, j int) bool {
	pairs := getPairs(rectangles, i, j)
	sortPairs(pairs)
	return canCut(pairs, 2)
}

func checkValidCuts(n int, rectangles [][]int) bool {
	return checkCut(n, rectangles, 0, 2) || checkCut(n, rectangles, 1, 3)
}
