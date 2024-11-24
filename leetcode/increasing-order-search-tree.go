/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getIncreasingBST(node *TreeNode) (*TreeNode, *TreeNode) {
	if node == nil {
		return nil, nil
	}
	leftHead, leftTail := getIncreasingBST(node.Left)
	rightHead, rightTail := getIncreasingBST(node.Right)
	node.Left, node.Right = nil, nil
	if leftHead == nil && leftTail == nil {
		leftHead, leftTail = node, node
	}
	if rightHead == nil && rightTail == nil {
		rightHead, rightTail = node, node
	}
	if leftTail != node {
		leftTail.Right = node
	}
	if node != rightHead {
		node.Right = rightHead
	}
	return leftHead, rightTail
}

func increasingBST(root *TreeNode) *TreeNode {
	head, _ := getIncreasingBST(root)
	return head
}
