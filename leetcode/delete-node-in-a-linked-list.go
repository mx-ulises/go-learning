/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(node *ListNode) {
	prev := node
	current := node.Next
	for current.Next != nil {
		prev.Val = current.Val
		prev = current
		current = current.Next
	}
	prev.Val = current.Val
	prev.Next = nil
}
