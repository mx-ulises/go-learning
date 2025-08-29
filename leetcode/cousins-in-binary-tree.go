package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getParentAndDepth(node *TreeNode, x int, parent *TreeNode, depth int) (*TreeNode, int) {
	if node == nil {
		return nil, -1
	} else if node.Val == x {
		return parent, depth
	}
	depth++
	lParent, lDepth := getParentAndDepth(node.Left, x, node, depth)
	if lDepth != -1 {
		return lParent, lDepth
	}
	return getParentAndDepth(node.Right, x, node, depth)
}

func isCousins(root *TreeNode, x int, y int) bool {
	xParent, xDepth := getParentAndDepth(root, x, nil, 0)
	yParent, yDepth := getParentAndDepth(root, y, nil, 0)
	return xParent != yParent && xDepth == yDepth
}
