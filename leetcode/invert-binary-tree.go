/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left, right := invertTree(root.Left), invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
