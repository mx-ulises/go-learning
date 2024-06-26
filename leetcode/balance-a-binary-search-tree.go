/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetInorderList(node *TreeNode, inorderList *[]*TreeNode) {
	if node == nil {
		return
	}
	GetInorderList(node.Left, inorderList)
	(*inorderList) = append(*inorderList, node)
	GetInorderList(node.Right, inorderList)
}

func BuildTree(inorderList *[]*TreeNode, start int, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (end + start) / 2
	(*inorderList)[mid].Left = BuildTree(inorderList, start, mid-1)
	(*inorderList)[mid].Right = BuildTree(inorderList, mid+1, end)
	return (*inorderList)[mid]
}

func balanceBST(root *TreeNode) *TreeNode {
	inorderList := []*TreeNode{}
	GetInorderList(root, &inorderList)
	return BuildTree(&inorderList, 0, len(inorderList)-1)
}
