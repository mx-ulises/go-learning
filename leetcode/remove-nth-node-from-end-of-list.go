/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func CountNodes(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a Dummy head to handle special case of removing head
	dummyHead := &ListNode{Next: head}
	// Get the position to be removed
	count := CountNodes(dummyHead) - n
	// Find the target node to be removed and the previous node
	prev, current := dummyHead, dummyHead
	for 0 < count {
		prev = current
		current = current.Next
		count--
	}
	// Remove the node from the list and return the original (or new) head
	prev.Next = current.Next
	return dummyHead.Next
}
