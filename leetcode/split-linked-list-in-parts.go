/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func Count(head *ListNode) int {
	current, count := head, 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := Count(head)
	mean := n / k
	extra := n % k
	output, current := make([]*ListNode, k), head
	for i := 0; i < k; i++ {
		if current != nil {
			output[i] = current
			elements := mean
			if 0 < extra {
				elements++
				extra--
			}
			var prev *ListNode
			for i := 0; i < elements; i++ {
				prev = current
				current = current.Next
			}
			prev.Next = nil
		}
	}
	return output
}
