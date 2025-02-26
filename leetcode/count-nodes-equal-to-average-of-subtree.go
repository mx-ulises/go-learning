/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getSubtrees(node *TreeNode) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	}
	sumLeft, countLeft, totalLeft := getSubtrees(node.Left)
	sumRight, countRight, totalRight := getSubtrees(node.Right)
	sum := sumLeft + sumRight + node.Val
	count := countLeft + countRight + 1
	total := totalLeft + totalRight
	if node.Val == sum/count {
		total++
	}
	return sum, count, total
}

func averageOfSubtree(root *TreeNode) int {
	_, _, total := getSubtrees(root)
	return total
}
