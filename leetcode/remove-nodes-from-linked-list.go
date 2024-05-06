/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNodes(head *ListNode) *ListNode {
	s := []*ListNode{}
	current := head
	for current != nil {
		for 0 < len(s) && s[len(s)-1].Val < current.Val {
			s = s[:len(s)-1]
		}
		s = append(s, current)
		current = current.Next
	}
	for i := 1; i < len(s); i++ {
		s[i-1].Next = s[i]
	}
	return s[0]
}
