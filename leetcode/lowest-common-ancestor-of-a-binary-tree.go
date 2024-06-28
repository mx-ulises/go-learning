/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetLowestCommonAncestor(node, p, q *TreeNode) (bool, bool, *TreeNode) {
	if node == nil {
		return false, false, nil
	}
	foundPLeft, foundQLeft, ancestor := GetLowestCommonAncestor(node.Left, p, q)
	if ancestor != nil {
		return true, true, ancestor
	}
	foundPRight, foundQRight, ancestor := GetLowestCommonAncestor(node.Right, p, q)
	if ancestor != nil {
		return true, true, ancestor
	}
	foundP := foundPLeft || foundPRight || node == p
	foundQ := foundQLeft || foundQRight || node == q
	if foundP && foundQ {
		ancestor = node
	}
	return foundP, foundQ, ancestor
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, output := GetLowestCommonAncestor(root, p, q)
	return output
}
