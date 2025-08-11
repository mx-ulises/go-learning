package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertToGreaterTree(node *TreeNode, total int) int {
	if node == nil {
		return total
	}
	node.Val += convertToGreaterTree(node.Right, total)
	return convertToGreaterTree(node.Left, node.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	_ = convertToGreaterTree(root, 0)
	return root
}
