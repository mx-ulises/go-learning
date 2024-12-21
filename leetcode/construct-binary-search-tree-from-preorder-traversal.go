/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insert(root *TreeNode, value int) {
	node := root
	for true {
		if value < node.Val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: value}
				break
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Val: value}
				break
			} else {
				node = node.Right
			}
		}
	}
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	for i := 1; i < len(preorder); i++ {
		insert(root, preorder[i])
	}
	return root
}
