/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func AddDigit(node *ListNode, l *ListNode) *ListNode {
	if l != nil {
		node.Val += l.Val
		return l.Next
	}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	carry := 0
	current := head
	for l1 != nil || l2 != nil || 0 < carry {
		succesor := new(ListNode)
		succesor.Val = carry
		l1 = AddDigit(succesor, l1)
		l2 = AddDigit(succesor, l2)
		carry = succesor.Val / 10
		succesor.Val %= 10
		current.Next = succesor
		current = succesor
	}
	return head.Next
}
