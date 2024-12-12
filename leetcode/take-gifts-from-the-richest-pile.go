// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	h := &IntHeap{}
	heap.Init(h)
	remainingGifts := 0
	for _, count := range gifts {
		heap.Push(h, count)
		remainingGifts += count
	}
	cache := map[int]int{}
	for i := 0; i < k; i++ {
		giftsInPile := heap.Pop(h).(int)
		remainingGifts -= giftsInPile
		remainingGiftsInPile, ok := cache[giftsInPile]
		if !ok {
			remainingGiftsInPile = int(math.Sqrt(float64(giftsInPile)))
			cache[giftsInPile] = remainingGiftsInPile
		}
		remainingGifts += remainingGiftsInPile
		heap.Push(h, remainingGiftsInPile)
	}
	return int64(remainingGifts)
}
