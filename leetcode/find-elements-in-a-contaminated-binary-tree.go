/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type FindElements struct {
	Elements map[int]bool
}

func RecoverAndAddElement(node *TreeNode, s []*TreeNode, val, n int) ([]*TreeNode, int) {
	if node != nil {
		node.Val = val
		s = append(s, node)
		n++
	}
	return s, n
}

func Constructor(root *TreeNode) FindElements {
	findElements := FindElements{Elements: map[int]bool{}}
	root.Val = 0
	s := []*TreeNode{root}
	n := 1
	for 0 < n {
		n--
		node := s[n]
		s = s[:n]
		findElements.Elements[node.Val] = true
		s, n = RecoverAndAddElement(node.Left, s, node.Val*2+1, n)
		s, n = RecoverAndAddElement(node.Right, s, node.Val*2+2, n)
	}
	return findElements
}

func (this *FindElements) Find(target int) bool {
	return this.Elements[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
