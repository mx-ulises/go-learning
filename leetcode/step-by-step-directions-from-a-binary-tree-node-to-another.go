/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GetCommonAncestor(node *TreeNode, start, end int) (*TreeNode, bool, bool) {
	if node == nil {
		return nil, false, false
	}
	candidateLeft, startFoundLeft, endFoundLeft := GetCommonAncestor(node.Left, start, end)
	if startFoundLeft && endFoundLeft {
		return candidateLeft, startFoundLeft, endFoundLeft
	}
	candidateRight, startFoundRight, endFoundRight := GetCommonAncestor(node.Right, start, end)
	if startFoundRight && endFoundRight {
		return candidateRight, startFoundRight, endFoundRight
	}
	var ancestor *TreeNode
	startFound := node.Val == start || startFoundLeft || startFoundRight
	endFound := node.Val == end || endFoundLeft || endFoundRight
	if startFound || endFound {
		ancestor = node
	}
	return ancestor, startFound, endFound
}

func GetPath(node *TreeNode, val int, path *[]byte, directions *[]byte) bool {
	if node == nil {
		return false
	}
	if node.Val == val {
		return true
	}
	n := len(*path)
	(*path) = append(*path, (*directions)[0])
	if GetPath(node.Left, val, path, directions) {
		return true
	}
	(*path)[n] = (*directions)[1]
	if GetPath(node.Right, val, path, directions) {
		return true
	}
	(*path) = (*path)[:n]
	return false
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	ancestor, _, _ := GetCommonAncestor(root, startValue, destValue)
	pathUp, pathDown := &[]byte{}, &[]byte{}
	GetPath(ancestor, startValue, pathUp, &[]byte{'U', 'U'})
	GetPath(ancestor, destValue, pathDown, &[]byte{'L', 'R'})
	return string(*pathUp) + string(*pathDown)
}
