/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
	Node   *TreeNode
	IsLeft bool
}

func IsLeave(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	s := []Pair{Pair{Node: root, IsLeft: false}}
	total := 0
	n := 1
	for 0 < n {
		n--
		pair := s[n]
		s = s[:n]
		if pair.Node == nil {
			continue
		} else if IsLeave(pair.Node) == false {
			s = append(s, Pair{Node: pair.Node.Left, IsLeft: true})
			s = append(s, Pair{Node: pair.Node.Right, IsLeft: false})
			n += 2
		} else if pair.IsLeft {
			total += pair.Node.Val
		}
	}
	return total
}

func DFS(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil && isLeft {
		return node.Val
	}
	return DFS(node.Left, true) + DFS(node.Right, false)
}

func sumOfLeftLeavesRecursive(root *TreeNode) int {
	return DFS(root, false)
}
