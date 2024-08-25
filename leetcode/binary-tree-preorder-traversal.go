/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func PreorderTraversal(node *TreeNode, output *[]int) {
	if node == nil {
		return
	}
	*output = append(*output, node.Val)
	PreorderTraversal(node.Left, output)
	PreorderTraversal(node.Right, output)
}

func preorderTraversal(root *TreeNode) []int {
	output := []int{}
	PreorderTraversal(root, &output)
	return output
}
