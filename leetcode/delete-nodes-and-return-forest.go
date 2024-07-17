/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BuildNodeMap(node *TreeNode, parent *TreeNode, parents *map[*TreeNode]*TreeNode, nodeMap *map[int]*TreeNode) {
	if node == nil {
		return
	}
	(*parents)[node] = parent
	(*nodeMap)[node.Val] = node
	BuildNodeMap(node.Left, node, parents, nodeMap)
	BuildNodeMap(node.Right, node, parents, nodeMap)
}

func DeleteFromParent(node *TreeNode, parents *map[*TreeNode]*TreeNode) {
	parent := (*parents)[node]
	if parent != nil {
		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
}

func DeleteFromChild(node *TreeNode, parents *map[*TreeNode]*TreeNode) {
	if node != nil {
		(*parents)[node] = nil
	}
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	parents := map[*TreeNode]*TreeNode{}
	nodeMap := map[int]*TreeNode{}
	BuildNodeMap(root, nil, &parents, &nodeMap)
	for _, val := range to_delete {
		node := nodeMap[val]
		DeleteFromParent(node, &parents)
		DeleteFromChild(node.Left, &parents)
		DeleteFromChild(node.Right, &parents)
		nodeMap[val] = nil
	}
	forest := []*TreeNode{}
	for _, node := range nodeMap {
		if node != nil && parents[node] == nil {
			forest = append(forest, node)
		}
	}
	return forest
}
