/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func fillReverseLists(node *TreeNode, reverseLists *map[int][]*TreeNode, level int) {
	if node == nil {
		return
	}
	if (level & 1) == 0 {
		(*reverseLists)[level] = append((*reverseLists)[level], node)
	}
	level++
	fillReverseLists(node.Left, reverseLists, level)
	fillReverseLists(node.Right, reverseLists, level)
}

func reverseNodeValues(reverseList *[]*TreeNode) {
	left := 0
	right := len(*reverseList) - 1
	for left < right {
		aux := (*reverseList)[left].Val
		(*reverseList)[left].Val = (*reverseList)[right].Val
		(*reverseList)[right].Val = aux
		left++
		right--
	}
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	reverseLists := map[int][]*TreeNode{}
	fillReverseLists(root, &reverseLists, 1)
	for _, reverseList := range reverseLists {
		reverseNodeValues(&reverseList)
	}
	return root
}
