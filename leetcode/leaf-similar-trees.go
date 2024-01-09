/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetLeaves(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	} else if (*node).Left == nil && (*node).Right == nil {
		(*leaves) = append(*leaves, (*node).Val)
	} else {
		GetLeaves((*node).Left, leaves)
		GetLeaves((*node).Right, leaves)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := make([]int, 0)
	GetLeaves(root1, &leaves1)
	leaves2 := make([]int, 0)
	GetLeaves(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}
