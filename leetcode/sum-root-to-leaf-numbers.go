/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetSumRootToLeafNum(node *TreeNode, current int) int {
	if node == nil {
		return 0
	}
	current = current*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return current
	} else {
		return GetSumRootToLeafNum(node.Left, current) + GetSumRootToLeafNum(node.Right, current)
	}
}

func sumNumbers(root *TreeNode) int {
	return GetSumRootToLeafNum(root, 0)
}
