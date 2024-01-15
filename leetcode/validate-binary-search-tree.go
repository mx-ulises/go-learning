/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsValidBST(node *TreeNode, start int64, end int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= start || end <= val {
		return false
	}
	return IsValidBST(node.Left, start, val) && IsValidBST(node.Right, val, end)
}

func isValidBST(root *TreeNode) bool {
	start := int64(-2147483649)
	end := int64(2147483648)
	return IsValidBST(root, start, end)
}
