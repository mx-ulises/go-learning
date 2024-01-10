/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	var minHead, minTail, maxHead, maxTail *ListNode
	current := head
	for current != nil {
		next := (*current).Next
		(*current).Next = nil
		if (*current).Val < x {
			if minHead == nil {
				minHead = current
				minTail = current
			} else {
				(*minTail).Next = current
				minTail = (*minTail).Next
			}
		} else {
			if maxHead == nil {
				maxHead = current
				maxTail = current
			} else {
				(*maxTail).Next = current
				maxTail = (*maxTail).Next
			}
		}
		current = next
	}
	if minTail != nil {
		(*minTail).Next = maxHead
	} else {
		minHead = maxHead
	}
	return minHead
}
