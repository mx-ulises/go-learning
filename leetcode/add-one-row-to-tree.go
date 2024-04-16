/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func AddRow(node *TreeNode, val int, depth int, isLeft bool) *TreeNode {
	depth--
	if depth == 0 {
		parent := &TreeNode{Val: val}
		if isLeft {
			parent.Left = node
		} else {
			parent.Right = node
		}
		node = parent
	} else if node != nil {
		node.Left = AddRow(node.Left, val, depth, true)
		node.Right = AddRow(node.Right, val, depth, false)
	}
	return node
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	return AddRow(root, val, depth, true)
}
