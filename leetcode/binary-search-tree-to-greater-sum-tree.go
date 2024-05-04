/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func ConvertToGst(node *TreeNode, current int) int {
	if node == nil {
		return current
	}
	node.Val += ConvertToGst(node.Right, current)
	return ConvertToGst(node.Left, node.Val)
}

func bstToGst(root *TreeNode) *TreeNode {
	ConvertToGst(root, 0)
	return root
}
