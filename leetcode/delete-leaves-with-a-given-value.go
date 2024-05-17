/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func RemoveLeafNodes(node *TreeNode, target int) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = RemoveLeafNodes(node.Left, target)
	node.Right = RemoveLeafNodes(node.Right, target)
	if node.Left == nil && node.Right == nil && node.Val == target {
		return nil
	}
	return node
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	return RemoveLeafNodes(root, target)
}
