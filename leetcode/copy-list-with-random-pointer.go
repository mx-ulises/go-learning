/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := map[*Node](*Node){}
	nodeMap[nil] = nil
	current := head
	for current != nil {
		nodeMap[current] = &Node{}
		current = current.Next
	}
	current = head
	for current != nil {
		nodeMap[current].Val = current.Val
		nodeMap[current].Next = nodeMap[current.Next]
		nodeMap[current].Random = nodeMap[current.Random]
		current = current.Next
	}
	return nodeMap[head]
}
