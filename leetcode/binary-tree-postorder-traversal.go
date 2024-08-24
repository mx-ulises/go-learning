/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Postorder(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}
	Postorder(node.Left, output)
	Postorder(node.Right, output)
	*output = append(*output, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	output := []int{}
	Postorder(root, &output)
	return output
}
