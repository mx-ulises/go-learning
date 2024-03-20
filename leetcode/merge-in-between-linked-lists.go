/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list2Tail := list2
	for list2Tail.Next != nil {
		list2Tail = list2Tail.Next
	}
	mockHead := &ListNode{Next: list1}
	prev, current := mockHead, list1
	i := 0
	for current != nil {
		if i == a {
			prev.Next = list2
		}
		if i == b {
			list2Tail.Next = current.Next
		}
		prev = current
		current = current.Next
		i++
	}
	return mockHead.Next
}
