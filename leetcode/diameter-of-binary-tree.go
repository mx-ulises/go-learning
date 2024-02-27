/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Diameter(node *TreeNode) (int, int) {
	if node == nil {
		return -1, -1
	}
	leftMaxDiameter, leftMaxHeight := Diameter(node.Left)
	rightMaxDiameter, rightMaxHeight := Diameter(node.Right)
	maxHeight := Max(leftMaxHeight, rightMaxHeight) + 1
	maxDiameter := Max(leftMaxHeight+rightMaxHeight+2, leftMaxDiameter)
	maxDiameter = Max(maxDiameter, rightMaxDiameter)
	return maxDiameter, maxHeight
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter, _ := Diameter(root)
	return maxDiameter
}
