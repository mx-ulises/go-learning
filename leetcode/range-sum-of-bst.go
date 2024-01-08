/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBSTRecursive(node *TreeNode, low int, high int, check bool) int {
	if check == false || node == nil {
		return 0
	}
	value := (*node).Val
	checkLeft := low <= value
	checkRight := value <= high
	total := rangeSumBSTRecursive((*node).Left, low, high, checkLeft)
	total += rangeSumBSTRecursive((*node).Right, low, high, checkRight)
	if checkLeft && checkRight {
		total += value
	}
	return total
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return rangeSumBSTRecursive(root, low, high, true)
}
