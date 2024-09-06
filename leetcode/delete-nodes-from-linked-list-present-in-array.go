/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap := map[int]bool{}
	for _, num := range nums {
		numsMap[num] = true
	}
	mockHead := &ListNode{Next: head}
	current := mockHead
	for current != nil {
		for current.Next != nil && numsMap[current.Next.Val] {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return mockHead.Next
}
