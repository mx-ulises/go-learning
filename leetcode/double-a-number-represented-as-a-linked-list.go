/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func Reverse(head *ListNode) *ListNode {
	var prev, current *ListNode = nil, head
	for current != nil {
		aux := current.Next
		current.Next = prev
		prev = current
		current = aux
	}
	return prev
}

func doubleIt(head *ListNode) *ListNode {
	head = Reverse(head)
	var prev, current *ListNode = nil, head
	carry := 0
	for current != nil {
		value := current.Val*2 + carry
		carry = value / 10
		current.Val = value % 10
		prev = current
		current = current.Next
	}
	if 0 < carry {
		prev.Next = &ListNode{Val: carry}
	}
	return Reverse(head)
}
