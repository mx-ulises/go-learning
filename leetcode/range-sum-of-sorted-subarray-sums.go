// Pair is a struct with two integers.
type Pair struct {
	CurrentSum int
	NextIndex  int
}

// Define a type that implements heap.Interface and holds Pairs.
type PairHeap []Pair

// Implement the heap.Interface methods.
func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].CurrentSum < h[j].CurrentSum }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Implement the Push method.
func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

// Implement the Pop method.
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rangeSum(nums []int, n int, left int, right int) int {
	h := &PairHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, Pair{nums[i], i + 1})
	}
	output := 0
	for i := 1; i <= right; i++ {
		pair := heap.Pop(h).(Pair)
		current, j := pair.CurrentSum, pair.NextIndex
		if left <= i {
			output += current
			output %= 1000000007
		}
		if pair.NextIndex < n {
			heap.Push(h, Pair{nums[j] + current, j + 1})
		}
	}
	return output
}
