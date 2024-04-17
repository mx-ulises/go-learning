/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func Compare(s1, s2 *[]int) *[]int {
	n, m := len(*s1), len(*s2)
	l := min(n, m)
	for i := 1; i <= l; i++ {
		j, k := n-i, m-i
		if (*s1)[j] < (*s2)[k] {
			return s1
		} else if (*s2)[k] < (*s1)[j] {
			return s2
		}
	}
	if l == m {
		return s2
	}
	return s1
}

func GetSmallerString(node *TreeNode, minimal, current *[]int) {
	if node == nil {
		return
	}
	(*current) = append(*current, node.Val)
	if node.Left == nil && node.Right == nil {
		minimalRef := Compare(minimal, current)
		newMinimal := make([]int, len(*minimalRef))
		copy(newMinimal, *minimalRef)
		(*minimal) = newMinimal
	} else {
		GetSmallerString(node.Left, minimal, current)
		GetSmallerString(node.Right, minimal, current)
	}
	(*current) = (*current)[:len(*current)-1]
}

func BuildString(minimal *[]int) string {
	output := []byte{}
	n := len(*minimal)
	for i := 0; i < n; i++ {
		output = append(output, 'a'+byte((*minimal)[n-i-1]))
	}
	return string(output)
}

func smallestFromLeaf(root *TreeNode) string {
	minimal := &[]int{26}
	current := &[]int{}
	GetSmallerString(root, minimal, current)
	return BuildString(minimal)
}
