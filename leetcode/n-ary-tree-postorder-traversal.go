/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func Postorder(node *Node, output *[]int) {
	if node == nil {
		return
	}
	for _, child := range node.Children {
		Postorder(child, output)
	}
	(*output) = append(*output, node.Val)
}

func postorder(root *Node) []int {
	output := []int{}
	Postorder(root, &output)
	return output
}
