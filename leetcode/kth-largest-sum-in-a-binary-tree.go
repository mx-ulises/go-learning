/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Define the struct
type Item struct {
	Sum   int
	Level int
}

// Define a type that implements heap.Interface and holds Items
type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].Sum == h[j].Sum {
		return h[i].Level < h[j].Level
	}
	return h[i].Sum < h[j].Sum
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

func GetLevelSum(node *TreeNode, level int, levels *map[int]int) {
	if node == nil {
		return
	}
	(*levels)[level] -= node.Val
	level++
	GetLevelSum(node.Left, level, levels)
	GetLevelSum(node.Right, level, levels)
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levels := map[int]int{}
	GetLevelSum(root, 1, &levels)
	h := &MinHeap{}
	for level, sum := range levels {
		heap.Push(h, Item{Sum: sum, Level: level})
	}
	for 0 < h.Len() && 1 < k {
		heap.Pop(h)
		k--
	}
	if h.Len() == 0 {
		return -1
	}
	return int64(-(*h)[0].Sum)
}
