/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkBalanceAndHeight(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	leftBalanced, leftHeight := checkBalanceAndHeight(node.Left)
	if leftBalanced == false {
		return false, 0
	}
	rightBalanced, rightHeight := checkBalanceAndHeight(node.Right)
	if rightBalanced == false {
		return false, 0
	}
	isBalanced := abs(leftHeight-rightHeight) < 2
	height := max(leftHeight, rightHeight) + 1
	return isBalanced, height
}

func isBalanced(root *TreeNode) bool {
	balanced, _ := checkBalanceAndHeight(root)
	return balanced
}
