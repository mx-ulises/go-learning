/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func AddNode(nodes *map[int]*TreeNode, val int) *TreeNode {
	if (*nodes)[val] == nil {
		(*nodes)[val] = &TreeNode{Val: val}
	}
	return (*nodes)[val]
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	hasParent := map[int]bool{}
	// Build Graph
	for _, description := range descriptions {
		parent, child, isLeft := description[0], description[1], description[2]
		parentNode := AddNode(&nodes, parent)
		childNode := AddNode(&nodes, child)
		hasParent[child] = true
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}
	// Find root
	var root *TreeNode
	for _, node := range nodes {
		if hasParent[node.Val] == false {
			root = node
			break
		}
	}
	return root
}
