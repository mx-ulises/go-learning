type DoubleLinkedNode struct {
	Val  int
	Prev *DoubleLinkedNode
	Succ *DoubleLinkedNode
}

func BuildDoubleLinkedList(n int) *DoubleLinkedNode {
	start := &DoubleLinkedNode{Val: 1, Prev: nil, Succ: nil}
	prev := start
	for i := 2; i <= n; i++ {
		current := &DoubleLinkedNode{Val: i, Prev: prev, Succ: nil}
		prev.Succ = current
		prev = current
	}
	prev.Succ = start
	start.Prev = prev
	return start
}

func MoveList(start *DoubleLinkedNode, k int) *DoubleLinkedNode {
	for i := 1; i < k; i++ {
		start = start.Succ
	}
	return start
}

func DeleteNode(node *DoubleLinkedNode) *DoubleLinkedNode {
	prev := node.Prev
	succ := node.Succ
	prev.Succ = succ
	succ.Prev = prev
	return succ
}

func findTheWinner(n int, k int) int {
	start := BuildDoubleLinkedList(n)
	for start != start.Succ {
		start = MoveList(start, k)
		start = DeleteNode(start)
	}
	return start.Val
}
