/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Evaluate(node *TreeNode) bool {
	var output bool
	switch node.Val {
	case 1:
		output = true
	case 2:
		output = Evaluate(node.Left) || Evaluate(node.Right)
	case 3:
		output = Evaluate(node.Left) && Evaluate(node.Right)
	default:
		output = false
	}
	return output
}

func evaluateTree(root *TreeNode) bool {
	return Evaluate(root)
}
