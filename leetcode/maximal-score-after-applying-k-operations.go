// MaxHeap represents a max-heap of integers.
type MaxHeap []int

// Implement heap.Interface and sort.Interface methods for MaxHeap.
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap condition
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	h := &MaxHeap{}
	for _, num := range nums {
		heap.Push(h, num)
	}
	score := 0
	for i := 0; i < k; i++ {
		num := heap.Pop(h).(int)
		score += num
		residual := num % 3
		num = num / 3
		if 0 < residual {
			num++
		}
		heap.Push(h, num)
	}
	return int64(score)
}
