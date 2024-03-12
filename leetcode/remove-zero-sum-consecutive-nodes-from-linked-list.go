/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	sumMap := map[int]*ListNode{0: dummyHead}
	sum := 0
	current := head
	for current != nil {
		sum += current.Val
		if sumMap[sum] != nil {
			removeSum := sum - current.Val
			for removeSum != sum {
				val := sumMap[removeSum].Val
				sumMap[removeSum] = nil
				removeSum -= val
			}
			sumMap[sum].Next = current.Next
		} else {
			sumMap[sum] = current
		}
		current = current.Next
	}
	return dummyHead.Next
}
