/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeNodes(head *ListNode) *ListNode {
	current := head
	pred := head
	for current.Next != nil {
		succ := current.Next
		if succ.Val == 0 {
			pred = current
			current = succ
		} else {
			current.Val += succ.Val
			current.Next = succ.Next
		}
	}
	pred.Next = nil
	return head
}
