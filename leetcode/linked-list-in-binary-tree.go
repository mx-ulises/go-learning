/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Check(listNode *ListNode, node *TreeNode) bool {
	if listNode == nil {
		return true
	}
	if node == nil || listNode.Val != node.Val {
		return false
	}
	return Check(listNode.Next, node.Left) || Check(listNode.Next, node.Right)
}

func isSubPath(head *ListNode, node *TreeNode) bool {
	if Check(head, node) {
		return true
	}
	if node == nil {
		return false
	}
	return isSubPath(head, node.Left) || isSubPath(head, node.Right)
}
