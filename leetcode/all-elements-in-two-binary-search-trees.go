/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getList(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	getList(node.Left, list)
	(*list) = append(*list, node.Val)
	getList(node.Right, list)
}

func mergeLists(list1, list2 []int) []int {
	list := make([]int, len(list1)+len(list2))
	i, j, k := 0, 0, 0
	for i < len(list) {
		if j == len(list1) {
			list[i] = list2[k]
			k++
		} else if k == len(list2) || list1[j] < list2[k] {
			list[i] = list1[j]
			j++
		} else {
			list[i] = list2[k]
			k++
		}
		i++
	}
	return list
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list1 := []int{}
	list2 := []int{}
	getList(root1, &list1)
	getList(root2, &list2)
	return mergeLists(list1, list2)
}
