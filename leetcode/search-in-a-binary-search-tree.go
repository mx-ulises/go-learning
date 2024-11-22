/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		switch {
		case val == root.Val:
			return root
		case val < root.Val:
			root = root.Left
		default:
			root = root.Right
		}
	}
	return nil
}
