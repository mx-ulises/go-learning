/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func ReverseListRecursive(prev, current *ListNode) *ListNode {
	if current == nil {
		return prev
	}
	newHead := ReverseListRecursive(current, current.Next)
	current.Next = prev
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	return ReverseListRecursive(nil, head)
}
