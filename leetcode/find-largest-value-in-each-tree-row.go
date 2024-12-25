/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Item struct {
	Node *TreeNode
	Row  int
}

func largestValues(root *TreeNode) []int {
	rows := make([]int, 0)
	q := []Item{{Node: root, Row: 0}}
	for 0 < len(q) {
		node, row := q[0].Node, q[0].Row
		q = q[1:]
		if node == nil {
			continue
		}
		if row == len(rows) {
			rows = append(rows, node.Val)
		} else {
			rows[row] = max(rows[row], node.Val)
		}
		row++
		q = append(q, Item{Node: node.Left, Row: row}, Item{Node: node.Right, Row: row})
	}
	return rows
}
