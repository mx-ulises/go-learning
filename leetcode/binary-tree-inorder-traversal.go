/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Traverse(node *TreeNode, output *[]int) {
	if node != nil {
		Traverse(node.Left, output)
		(*output) = append((*output), node.Val)
		Traverse(node.Right, output)
	}
}

func inorderTraversal(root *TreeNode) []int {
	output := []int{}
	Traverse(root, &output)
	return output
}
