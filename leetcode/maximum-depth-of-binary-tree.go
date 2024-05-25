/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(GetMaxDepth(node.Left), GetMaxDepth(node.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return GetMaxDepth(root)
}
