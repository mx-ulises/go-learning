type Pair struct {
	Index int
	Num   int
}

func addPairs(output []int, currentPairs []Pair) {
	n := len(currentPairs)
	indexes := make([]int, n)
	for i := 0; i < n; i++ {
		indexes[i] = currentPairs[i].Index
	}
	sort.Ints(indexes)
	for i := 0; i < n; i++ {
		index, num := indexes[i], currentPairs[i].Num
		output[index] = num
	}
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	pairs := make([]Pair, n)
	for i, num := range nums {
		pairs[i] = Pair{Index: i, Num: num}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Num < pairs[j].Num
	})
	currentPairs := []Pair{}
	output := make([]int, n)
	for _, pair := range pairs {
		l := len(currentPairs)
		if l == 0 || (pair.Num-currentPairs[l-1].Num) <= limit {
			currentPairs = append(currentPairs, pair)
		} else {
			addPairs(output, currentPairs)
			currentPairs = []Pair{pair}
		}
	}
	addPairs(output, currentPairs)
	return output
}
