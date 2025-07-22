package leetcode

import "container/heap"

// IntHeap defines a min-heap using a slice of ints
type IntHeap []int

// Len returns the number of elements in the heap
func (h IntHeap) Len() int { return len(h) }

// Less defines the heap property — for min-heap, parent < children
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap exchanges elements at indexes i and j
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds a new element to the heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element (root)
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	min := old[n-1]
	*h = old[0 : n-1]
	return min
}

func trimMean(arr []int) float64 {
	heapSize := len(arr) / 20
	minHeap, maxHeap := &IntHeap{}, &IntHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	sum := 0
	for _, num := range arr {
		heap.Push(minHeap, num)
		heap.Push(maxHeap, -num)
		if heapSize < minHeap.Len() {
			heap.Pop(minHeap)
			heap.Pop(maxHeap)
		}
		sum += num
	}
	for 0 < minHeap.Len() {
		sum -= heap.Pop(minHeap).(int)
		sum += heap.Pop(maxHeap).(int)
	}
	return float64(sum) / float64(heapSize*18)
}
