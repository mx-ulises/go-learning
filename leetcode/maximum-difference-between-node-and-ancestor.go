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

func maxAncestorDiffTraversal(node *TreeNode, maximal, minimal int) int {
	if node == nil {
		return 0
	}
	diff := max(abs(node.Val-maximal), abs(node.Val-minimal))
	maximal, minimal = max(maximal, node.Val), min(minimal, node.Val)
	leftDiff := maxAncestorDiffTraversal(node.Left, maximal, minimal)
	rightDiff := maxAncestorDiffTraversal(node.Right, maximal, minimal)
	return max(diff, leftDiff, rightDiff)
}

func maxAncestorDiff(root *TreeNode) int {
	return maxAncestorDiffTraversal(root, root.Val, root.Val)
}
