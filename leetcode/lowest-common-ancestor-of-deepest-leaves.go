/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getLca(node *TreeNode, level int) (*TreeNode, int) {
	if node == nil {
		return nil, level
	}
	level++
	left, leftLevel := getLca(node.Left, level)
	right, rightLevel := getLca(node.Right, level)
	if leftLevel < rightLevel {
		return right, rightLevel
	}
	if rightLevel < leftLevel {
		return left, leftLevel
	}
	return node, leftLevel
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	lca, _ := getLca(root, 0)
	return lca
}
