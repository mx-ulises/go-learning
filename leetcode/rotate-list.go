/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func GetListSize(head *ListNode) int {
	size := 0
	current := head
	for current != nil {
		current = current.Next
		size += 1
	}
	return size
}

func MakeCircular(head *ListNode) *ListNode {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = head
	return current
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	size := GetListSize(head)
	prev := MakeCircular(head)
	moves := size - (k % size)
	for i := 0; i < moves; i++ {
		prev = prev.Next
		head = head.Next
	}
	prev.Next = nil
	return head
}

