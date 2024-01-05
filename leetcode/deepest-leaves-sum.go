/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodePair struct {
	Level int
	Node  *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	maximalLevel := 0
	currentSum := 0
	s := make([]TreeNodePair, 0)
	s = append(s, TreeNodePair{0, root})
	for 0 < len(s) {
		pair := s[len(s)-1]
		s = s[:len(s)-1]
		if pair.Node == nil {
			continue
		}
		if maximalLevel < pair.Level {
			currentSum = 0
			maximalLevel = pair.Level
		}
		if maximalLevel == pair.Level {
			currentSum += pair.Node.Val
		}
		s = append(s, TreeNodePair{pair.Level + 1, pair.Node.Left})
		s = append(s, TreeNodePair{pair.Level + 1, pair.Node.Right})
	}
	return currentSum
}
