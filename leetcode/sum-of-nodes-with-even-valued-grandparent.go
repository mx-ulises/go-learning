/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func TraverseSumEvenGP(node *TreeNode, parentValue, grandparentValue int) int {
	if node == nil {
		return 0
	}
	output := 0
	if (grandparentValue & 1) == 0 {
		output += node.Val
	}
	grandparentValue = parentValue
	parentValue = node.Val
	output += TraverseSumEvenGP(node.Left, parentValue, grandparentValue)
	output += TraverseSumEvenGP(node.Right, parentValue, grandparentValue)
	return output
}

func sumEvenGrandparent(root *TreeNode) int {
	return TraverseSumEvenGP(root, 1, 1)
}
