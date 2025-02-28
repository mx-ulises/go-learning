/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverFromPreorder(traversal string) *TreeNode {
	i, n := 0, len(traversal)
	s := []*TreeNode{}
	level, val := 0, 0
	for i < n {
		for i < n && traversal[i] == '-' {
			level++
			i++
		}
		for i < n && traversal[i] != '-' {
			val = val*10 + int(traversal[i]-'0')
			i++
		}
		s = s[:level]
		child := &TreeNode{Val: val}
		s = append(s, child)
		if 0 < level {
			parent := s[level-1]
			if parent.Left == nil {
				parent.Left = child
			} else {
				parent.Right = child
			}
		}
		level, val = 0, 0
	}
	return s[0]
}
