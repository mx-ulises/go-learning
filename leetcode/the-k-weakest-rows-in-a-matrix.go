type Item struct {
	Strength int
	Index    int
}

type MinHeap []Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	if pq[i].Strength == pq[j].Strength {
		return pq[i].Index > pq[j].Index
	}
	return pq[i].Strength > pq[j].Strength
}

func (pq MinHeap) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *MinHeap) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func (pq *MinHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func getStrength(row []int) int {
	s := 0
	for _, x := range row {
		s += x
	}
	return s
}

func kWeakestRows(mat [][]int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)
	for i, row := range mat {
		s := getStrength(row)
		heap.Push(h, Item{Strength: s, Index: i})
		if k < h.Len() {
			heap.Pop(h)
		}
	}
	output := make([]int, h.Len())
	i := h.Len() - 1
	for 0 <= i {
		item := heap.Pop(h).(Item)
		output[i] = item.Index
		i--
	}
	return output
}
