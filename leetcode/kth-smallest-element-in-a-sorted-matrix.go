type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			heap.Push(h, matrix[i][j])
			if k < len(*h) {
				heap.Pop(h)
			}
		}
	}
	return (*h)[0]
}
