/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		root1 = &TreeNode{}
	}
	if root2 == nil {
		root2 = &TreeNode{}
	}
	root := &TreeNode{}
	root.Val = root1.Val + root2.Val
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
