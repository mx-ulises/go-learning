/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func pairSum(head *ListNode) int {
	current, runner := head, head
	for runner != nil {
		current = current.Next
		runner = runner.Next.Next
	}
	var tail *ListNode
	for current != nil {
		succ := current.Next
		current.Next = tail
		tail = current
		current = succ
	}
	maxPairSum := 0
	for tail != nil {
		maxPairSum = max(maxPairSum, head.Val+tail.Val)
		head = head.Next
		tail = tail.Next
	}
	return maxPairSum
}
