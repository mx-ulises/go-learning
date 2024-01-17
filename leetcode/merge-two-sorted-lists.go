/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedHead := &ListNode{}
	prev := mergedHead
	for list1 != nil || list2 != nil {
		if (list1 == nil) || (list2 != nil && list2.Val < list1.Val) {
			prev.Next = list2
			list2 = list2.Next
		} else {
			prev.Next = list1
			list1 = list1.Next
		}
		prev = prev.Next
	}
	return mergedHead.Next
}
