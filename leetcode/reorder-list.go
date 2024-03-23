/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// get the middle point of the list
	current1 := head
	current2 := head
	for current2 != nil && current2.Next != nil {
		current1 = current1.Next
		current2 = current2.Next.Next
	}
	// If it is odd number, push one more to the left
	if current2 != nil {
		current1 = current1.Next
	}
	// Reverse list from the middle
	current2 = nil
	for current1 != nil {
		next := current1.Next
		current1.Next = current2
		current2 = current1
		current1 = next
	}
	// Merge head and reversed list
	current1 = head
	for current2 != nil {
		next1 := current1.Next
		next2 := current2.Next
		current1.Next = current2
		current2.Next = next1
		current1 = next1
		current2 = next2
	}
	// If last element in the head list hasn't been proccesed, tail it to nil
	if current1 != nil {
		current1.Next = nil
	}
}
