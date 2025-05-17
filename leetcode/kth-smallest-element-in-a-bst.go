/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findNode(node *TreeNode, k int) (int, int) {
	if node == nil {
		return k, -1
	}
	k, val := findNode(node.Left, k)
	if val != -1 {
		return k, val
	}
	k--
	if k == 0 {
		return k, node.Val
	}
	k, val = findNode(node.Right, k)
	if val != -1 {
		return k, val
	}
	return k, -1
}

func kthSmallest(root *TreeNode, k int) int {
	_, val := findNode(root, k)
	return val
}
