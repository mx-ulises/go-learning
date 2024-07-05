/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isCriticalPoint(prev int, current int, next int) bool {
	if prev < current && next < current {
		return true
	}
	if current < prev && current < next {
		return true
	}
	return false
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	prev, current := head, head.Next
	firstCriticalPoint, prevCriticalPoint := -1, -1
	minDistance, maxDistance := 100001, -1
	i := 1
	for current.Next != nil {
		if isCriticalPoint(prev.Val, current.Val, current.Next.Val) {
			if firstCriticalPoint == -1 {
				firstCriticalPoint = i
			} else {
				maxDistance = i - firstCriticalPoint
				minDistance = min(minDistance, i-prevCriticalPoint)
			}
			prevCriticalPoint = i
		}
		i++
		prev = current
		current = current.Next
	}
	if minDistance == 100001 {
		minDistance = -1
	}
	return []int{minDistance, maxDistance}
}
