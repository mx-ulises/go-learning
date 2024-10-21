// Define the struct
type Item struct {
	Val   int
	Index int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Val == h[j].Val {
		return h[i].Index < h[j].Index
	}
	return h[i].Val < h[j].Val
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func getFinalState(nums []int, k int, multiplier int) []int {
	h := &MinHeap{}
	for i, num := range nums {
		heap.Push(h, Item{Val: num, Index: i})
	}
	for i := 0; i < k; i++ {
		item := heap.Pop(h).(Item)
		item.Val *= multiplier
		heap.Push(h, item)
	}
	for _, item := range *h {
		nums[item.Index] = item.Val
	}
	return nums
}
