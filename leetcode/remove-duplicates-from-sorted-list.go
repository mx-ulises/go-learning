/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	mockHead := &ListNode{Val: -101, Next: head}
	current := head
	prev := mockHead
	for current != nil {
		for current != nil && prev.Val == current.Val {
			prev.Next = current.Next
			current = current.Next
		}
		prev = prev.Next
		if current != nil {
			current = current.Next
		}
	}
	return mockHead.Next
}
