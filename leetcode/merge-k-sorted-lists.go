/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap MinHeap
	heap.Init(&minHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}
	mockHead := &ListNode{Next: nil}
	current := mockHead
	for minHeap.Len() > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		current.Next = node
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
		current = node
	}
	return mockHead.Next
}
