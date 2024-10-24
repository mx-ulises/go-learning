/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func CheckEquiv(node1 *TreeNode, node2 *TreeNode) bool {
	// Both nodes are null, they are equivalent
	if node1 == nil && node2 == nil {
		return true
	}
	// One of the nodes is null, or their values are diff, they are not equivalent
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}
	// Check subtrees without flip
	if CheckEquiv(node1.Left, node2.Left) && CheckEquiv(node1.Right, node2.Right) {
		return true
	}
	// Check subtrees with flip
	return CheckEquiv(node1.Right, node2.Left) && CheckEquiv(node1.Left, node2.Right)
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return CheckFlipEquivalence(root1, root2)
}
