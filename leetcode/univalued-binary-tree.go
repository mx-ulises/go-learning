/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func testSubtree(node *TreeNode, parentVal int) bool {
	return isUnivalTree(node) && (node == nil || node.Val == parentVal)
}

func isUnivalTree(node *TreeNode) bool {
	if node == nil {
		return true
	}
	return testSubtree(node.Left, node.Val) && testSubtree(node.Right, node.Val)
}
