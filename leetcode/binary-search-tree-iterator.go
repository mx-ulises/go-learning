/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.addSubtree(root)
	return iterator
}

func (this *BSTIterator) addSubtree(node *TreeNode) {
	for node != nil {
		this.Stack = append(this.Stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.addSubtree(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return 0 < len(this.Stack)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
 