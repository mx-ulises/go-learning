/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetLevels(node *TreeNode, levels *map[int]int, level int) {
	if node == nil {
		return
	}
	(*levels)[level] += node.Val
	level++
	GetLevels(node.Left, levels, level)
	GetLevels(node.Right, levels, level)
}

func GetNodeValue(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func SetNodeValue(node *TreeNode, value int) {
	if node == nil {
		return
	}
	node.Val = value
}

func ReplaceChildValue(node *TreeNode, levels *map[int]int, level int) {
	if node == nil {
		return
	}
	level++
	value := (*levels)[level] - GetNodeValue(node.Left) - GetNodeValue(node.Right)
	SetNodeValue(node.Left, value)
	SetNodeValue(node.Right, value)
	ReplaceChildValue(node.Left, levels, level)
	ReplaceChildValue(node.Right, levels, level)
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	levels := map[int]int{}
	GetLevels(root, &levels, 0)
	ReplaceChildValue(root, &levels, 0)
	root.Val = 0
	return root
}
